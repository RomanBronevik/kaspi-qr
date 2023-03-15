package main

import (
	"context"
	"log"

	"kaspi-qr/config"
	"kaspi-qr/internal/adapters/provider/kaspi"
	"kaspi-qr/internal/adapters/repo/pg"
	"kaspi-qr/internal/adapters/server"
	"kaspi-qr/internal/adapters/server/rest"
	"kaspi-qr/internal/domain/core"
	"kaspi-qr/internal/domain/usecases"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

func main() {
	app := struct {
		repo  *pg.St
		core  *core.St
		ucs   *usecases.St
		srv   *server.St
		kaspi *kaspi.St
	}{}

	conf := config.Load()

	if err := initConfig(); err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		log.Fatalf("error loading env variables: %s", err.Error())
	}

	postgreSQLClient, err := pg.NewClient(context.TODO())
	if err != nil {
		return
	}

	srv := new(server.St)

	app.repo = pg.NewRepository(postgreSQLClient)
	app.ucs = usecases.New(app.repo)
	app.srv = srv

	app.kaspi = kaspi.New()

	app.core = core.New(app.repo, app.kaspi)
	app.ucs.SetCore(app.core)

	handlers := rest.NewHandler(app.ucs, app.kaspi)

	srv.Run(viper.GetString("port"), handlers.InitRoutes())
}
