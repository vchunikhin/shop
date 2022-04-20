package main

import (
	"github.com/spf13/viper"
	"local/shop/backend/pkg/handler"
	"local/shop/backend/pkg/repository"
	"local/shop/backend/pkg/service"
	"log"
)

const (
	configPath = "configs"
	configName = "config"
	portKey = "port"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}
	rep := repository.NewRepository()
	serv := service.NewService(rep)
	handlers := handler.NewHandler(serv)
	srv := new(Server)
	if err := srv.Run(viper.GetString(portKey), handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath(configPath)
	viper.SetConfigName(configName)
	return viper.ReadInConfig()
}


