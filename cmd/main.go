package main

import (
	"kaspi-qr/internal/adapters/server/rest"
	"os"
	"os/signal"
	"syscall"
	"time"

	"kaspi-qr/config"
	dbPg "kaspi-qr/internal/adapters/db/pg"
	"kaspi-qr/internal/adapters/logger/zap"
	"kaspi-qr/internal/adapters/provider/kaspi"
	repoPg "kaspi-qr/internal/adapters/repo/pg"
	"kaspi-qr/internal/adapters/server"
	"kaspi-qr/internal/domain/core"
	"kaspi-qr/internal/domain/usecases"
)

func main() {
	var err error

	app := struct {
		lg    *zap.St
		db    *dbPg.St
		repo  *repoPg.St
		core  *core.St
		ucs   *usecases.St
		srv   *server.St
		kaspi *kaspi.St
	}{}

	// load config
	conf := config.Load()

	// logger
	app.lg = zap.New(conf.LogLevel, conf.Debug)

	// db
	app.db, err = dbPg.New(conf.Debug, app.lg, dbPg.OptionsSt{
		Dsn: conf.PgDsn,
	})
	if err != nil {
		app.lg.Fatal(err)
	}

	// kaspi
	app.kaspi, err = kaspi.New(app.lg, conf.KaspiApiUrl, conf.CertPath, conf.CertPsw)
	if err != nil {
		app.lg.Fatal(err)
	}

	// repo
	app.repo = repoPg.New(app.lg, app.db)

	// core
	app.core = core.New(app.lg, app.repo, app.kaspi, conf.QrUrlTemplate)

	// usecases
	app.ucs = usecases.New(app.lg, app.db, app.core)

	// START

	app.lg.Infow("Starting")

	app.srv = server.Start(
		app.lg,
		conf.HttpListen,
		rest.GetHandler(app.lg, app.ucs, conf.HttpCors),
	)

	var exitCode int

	select {
	case <-stopSignal():
	case <-app.srv.Wait():
		exitCode = 1
	}

	// STOP

	app.lg.Infow("Shutting down...")

	if !app.srv.Shutdown(20 * time.Second) {
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
