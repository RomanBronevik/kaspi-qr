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
	notifierHttp "kaspi-qr/internal/adapters/notifier/http"
	"kaspi-qr/internal/adapters/provider/kaspi"
	repoPg "kaspi-qr/internal/adapters/repo/pg"
	"kaspi-qr/internal/adapters/server"
	"kaspi-qr/internal/domain/core"
	"kaspi-qr/internal/domain/usecases"

	dopDbPg "github.com/rendau/dop/adapters/db/pg"
)

func main() {
	var err error

	app := struct {
		lg       *zap.St
		db       *dopDbPg.St
		dbRaw    *dbPg.St
		repo     *repoPg.St
		core     *core.St
		ucs      *usecases.St
		srv      *server.St
		kaspi    *kaspi.St
		notifier *notifierHttp.St
	}{}

	// load config
	conf := config.Load()

	// logger
	app.lg = zap.New(conf.LogLevel, conf.Debug)

	// db
	app.db, err = dopDbPg.New(conf.Debug, app.lg, dopDbPg.OptionsSt{
		Dsn: conf.PgDsn,
	})
	if err != nil {
		app.lg.Fatal(err)
	}

	// dbRaw
	app.dbRaw, err = dbPg.New(conf.Debug, app.lg, dbPg.OptionsSt{
		Dsn: conf.PgDsn,
	})
	if err != nil {
		app.lg.Fatal(err)
	}

	// repo
	app.repo = repoPg.New(app.lg, app.db, app.dbRaw)

	// kaspi
	app.kaspi, err = kaspi.New(app.lg, conf.KaspiApiUrl, conf.CertPath, conf.CertPsw)
	if err != nil {
		app.lg.Fatal(err)
	}

	// notifier
	app.notifier = notifierHttp.New(app.lg)

	// core
	app.core = core.New(
		app.lg,
		app.repo,
		app.kaspi,
		app.notifier,
		conf.QrUrlTemplate,
	)

	// usecases
	app.ucs = usecases.New(app.lg, app.dbRaw, app.core)

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
