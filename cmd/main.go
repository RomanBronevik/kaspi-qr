package main

import (
	"fmt"
	"github.com/spf13/viper"
	kaspi_qr "kaspi-qr"
	"kaspi-qr/pkg/handler"
	"kaspi-qr/pkg/repository"
	"kaspi-qr/pkg/service"
	"log"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}
	fmt.Print("asdasdasd")
	//asdasd123123123123123
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
