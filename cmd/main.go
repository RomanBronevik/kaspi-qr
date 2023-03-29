package main

import (
	"os"
	"os/signal"
	"syscall"
	"time"

	"kaspi-qr/config"
	dbPg "kaspi-qr/internal/adapters/db/pg"
	"kaspi-qr/internal/adapters/logger/zap"
	"kaspi-qr/internal/adapters/provider/kaspi"
	repoPg "kaspi-qr/internal/adapters/repo/pg"
	"kaspi-qr/internal/adapters/server/rest"
	"kaspi-qr/internal/adapters/server/rest/h_gin"
	"kaspi-qr/internal/domain/core"
	"kaspi-qr/internal/domain/usecases"

	_ "github.com/lib/pq"
)

func main() {
	var err error

	app := struct {
		lg          *zap.St
		db          *dbPg.St
		repo        *repoPg.St
		core        *core.St
		ucs         *usecases.St
		restHandler *h_gin.Handler
		restSrv     *rest.St
		kaspi       *kaspi.St
	}{}

	// load config
	conf := config.Load()

	// logger
	app.lg = zap.New(conf.LogLevel, conf.Debug)

	// db
	app.db, err = dbPg.New(app.lg, dbPg.OptionsSt{
		Dsn: conf.PgDsn,
	})
	if err != nil {
		app.lg.Fatal(err)
	}

	// kaspi
	app.kaspi = kaspi.New(app.lg, conf.KaspiApiUrl, conf.CertPath, conf.CertPassword)

	// repo
	app.repo = repoPg.New(app.lg, app.db)

	// core
	app.core = core.New(app.repo, app.kaspi)

	// usecases
	app.ucs = usecases.New(app.lg, app.db, app.core)

	// rest handler
	app.restHandler = h_gin.NewHandler(app.lg, app.ucs)

	// START

	app.lg.Infow("Starting")

	app.restSrv = rest.Start(
		app.lg,
		conf.HttpListen,
		app.restHandler.InitRoutes(conf.HttpCors),
	)

	var exitCode int

	select {
	case <-stopSignal():
	case <-app.restSrv.Wait():
		exitCode = 1
	}

	// STOP

	app.lg.Infow("Shutting down...")

	if !app.restSrv.Shutdown(20 * time.Second) {
		exitCode = 1
	}

	app.lg.Infow("Wait routines...")

	app.core.WaitJobs()

	app.lg.Infow("Exit")

	os.Exit(exitCode)
}

func stopSignal() <-chan os.Signal {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	return ch
}
