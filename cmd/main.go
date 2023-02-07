package main

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	kaspi_qr "kaspi-qr"
	"kaspi-qr/internal/handler"
	organization2 "kaspi-qr/internal/organization"
	organization "kaspi-qr/internal/organization/db"
	"kaspi-qr/pkg/repository"
	"kaspi-qr/pkg/service"
	"log"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}
	//fmt.Print("hello world")

	if err := godotenv.Load(); err != nil {
		log.Fatalf("error loading env variables: %s", err.Error())
	}

	postgreSQLClient, err := repository.NewClient(context.TODO())
	if err != nil {
		return
	}

	repo := organization.NewRepository(postgreSQLClient)

	newOrg := organization2.Organization{
		Bin: 1337235,
	}

	err = repo.Create(context.TODO(), &newOrg)
	if err != nil {
		fmt.Println(err.Error())
	}

	all, err := repo.FindAll(context.TODO())
	if err != nil {
		return
	}

	for _, ath := range all {
		fmt.Println(ath)
	}
	//db, err :=
	//db, err := repository.NewPostgresDB(repository.Config{
	//	Host:     viper.GetString("db.host"),
	//	Port:     viper.GetString("db.port"),
	//	Username: viper.GetString("db.username"),
	//	Password: os.Getenv("DB_PASSWORD"),
	//	DBName:   viper.GetString("db.dbname"),
	//	SSLMode:  viper.GetString("db.sslmode"),
	//})

	//if err != nil {
	//	log.Fatalf("failed to initialize DB: %s", err.Error())
	//}
	//
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(kaspi_qr.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig() // считывает значение конфиги и записывает их во внутренний объект вайпера и возвращает ошибку
}
