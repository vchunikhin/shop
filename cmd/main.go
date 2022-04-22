package main

import (
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"local/shop/backend/pkg/handler"
	"local/shop/backend/pkg/repository"
	"local/shop/backend/pkg/service"
	"log"
	"os"
)

const (
	configPath = "configs"
	configName = "config"
	serverPortKey = "port"
	dbHostKey = "db.host"
	dbPortKey = "db.port"
	dbUsernameKey = "db.username"
	dbPasswordKey = "DB_PASSWORD"
	dbName = "db.name"
	dbSSLMode = "db.sslmode"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		log.Fatalf("error loading env variables: %s", err.Error())
	}

	db, err := initDB()
	if err != nil {
		log.Fatalf("failed to initialize db: %s", err.Error())
	}

	rep := repository.NewRepository(db)
	serv := service.NewService(rep)
	handlers := handler.NewHandler(serv)
	srv := new(Server)
	if err := srv.Run(viper.GetString(serverPortKey), handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath(configPath)
	viper.SetConfigName(configName)
	return viper.ReadInConfig()
}

func initDB() (*sqlx.DB, error) {
	return repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString(dbHostKey),
		Port:     viper.GetString(dbPortKey),
		Username: viper.GetString(dbUsernameKey),
		Password: os.Getenv(dbPasswordKey),
		DBName:   viper.GetString(dbName),
		SSLMode:  viper.GetString(dbSSLMode),
	})
}


