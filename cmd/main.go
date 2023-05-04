package main

import (
	notifierHttp "kaspi-qr/internal/adapters/notifier/http"
	"kaspi-qr/internal/adapters/provider/kaspi"
	repoPg "kaspi-qr/internal/adapters/repo/pg"
	"kaspi-qr/internal/adapters/server/rest"
	"kaspi-qr/internal/domain/core"
	"kaspi-qr/internal/domain/usecases"
	"os"
	"time"

	dopDbPg "github.com/rendau/dop/adapters/db/pg"
	dopLoggerZap "github.com/rendau/dop/adapters/logger/zap"
	dopServerHttps "github.com/rendau/dop/adapters/server/https"
	"github.com/rendau/dop/dopTools"
	"github.com/swaggo/swag/example/basic/docs"
)

func main() {
	var err error

	app := struct {
		lg         *dopLoggerZap.St
		db         *dopDbPg.St
		repo       *repoPg.St
		core       *core.St
		ucs        *usecases.St
		kaspi      *kaspi.St
		notifier   *notifierHttp.St
		restApiSrv *dopServerHttps.St
	}{}

	// load config
	conf := ConfLoad()

	// logger
	app.lg = dopLoggerZap.New(conf.LogLevel, conf.Debug)

	app.db, err = dopDbPg.New(conf.Debug, app.lg, dopDbPg.OptionsSt{
		Dsn:      conf.PgDsn,
		Timezone: "Asia/Almaty",
	})
	if err != nil {
		app.lg.Fatal(err)
	}

	// repo
	app.repo = repoPg.New(app.lg, app.db)

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
	app.ucs = usecases.New(app.lg, app.db, app.core)

	docs.SwaggerInfo.Host = conf.SwagHost
	docs.SwaggerInfo.BasePath = conf.SwagBasePath
	docs.SwaggerInfo.Schemes = []string{conf.SwagSchema}
	docs.SwaggerInfo.Title = "Kaspi-QR"

	// START

	app.lg.Infow("Starting")

	app.restApiSrv = dopServerHttps.Start(
		conf.HttpListen,
		rest.GetHandler(
			app.lg,
			app.ucs,
			conf.HttpCors,
		),
		app.lg,
	)

	app.core.Start()

	// LISTEN

	var exitCode int

	select {
	case <-dopTools.StopSignal():
	case <-app.restApiSrv.Wait():
		exitCode = 1
	}

	// STOP

	app.lg.Infow("Shutting down...")

	if !app.restApiSrv.Shutdown(20 * time.Second) {
		exitCode = 1
	}

	app.lg.Infow("Wait routines...")

	app.core.StopAndWaitJobs()

	app.lg.Infow("Exit")

	os.Exit(exitCode)
}
