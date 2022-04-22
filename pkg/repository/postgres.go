package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
)

const (
	dbDriverKey = "db.driver"
	dbSchemaKey = "DB_SCHEMA"

	categoriesTable = "categories"
	itemsTable = "items"
	ordersTable = "orders"
	productsTable = "products"
	usersTable = "users"
)

type Config struct {
	Host string
	Port string
	Username string
	Password string
	DBName string
	SSLMode string
}

func NewPostgresDB(cfg Config) (*sqlx.DB, error) {
	db, err := sqlx.Open(viper.GetString(dbDriverKey), fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode))

	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
