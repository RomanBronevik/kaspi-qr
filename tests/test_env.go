package tests

import (
	"context"
	"kaspi-qr/internal/adapters/provider/kaspi"
	"kaspi-qr/internal/adapters/repo/pg"
	"kaspi-qr/internal/domain/core"
	"kaspi-qr/internal/domain/usecases"
	"log"

	"github.com/spf13/viper"
)

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig() // считывает значение конфиги и записывает их во внутренний объект вайпера и возвращает ошибку
}

func ViperAndOsConfig() {
	if err := initConfig(); err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}

	postgreSQLClient, err := pg.NewClient(context.TODO())
	if err != nil {
		return
	}

	app.repo = pg.NewRepository(postgreSQLClient)
	app.usc = usecases.New(app.repo)
	app.kaspi = kaspi.New()
	app.cr = core.New(app.repo, app.kaspi)
	app.usc.SetCore(app.cr)
	//viper.SetConfigType("yaml")
	//viper.SetConfigFile("config.yaml")
	//viper.ReadInConfig()
	//
	//viper.Set("kaspiUrl", "https://mtokentest.kaspi.kz:8545/r3/v01/")
	//viper.Set("port", "8000")
	//viper.Set("testBin", "160640004075")
	//
	//viper.Set("db.username", "postgres")
	//viper.Set("db.host", "localhost")
	//viper.Set("db.port", "5436")
	//viper.Set("dbname", "postgres")
	//viper.Set("maxAttempts", 5)
	//
	//viper.WriteConfig()
	//
	//os.Setenv("CERTIFICATE_PATH", "./configs/certfile.pfx")
	//os.Setenv("CERTIFICATE_PASSWORD", "asd123456")
	//os.Setenv("DB_PASSWORD", "asd123456")

}
