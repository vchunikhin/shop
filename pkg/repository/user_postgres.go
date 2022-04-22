package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	shop "local/shop/backend"
	"os"
)

type UserPostgres struct {
	db *sqlx.DB
}

func NewUserPostgres(db *sqlx.DB) *UserPostgres {
	return &UserPostgres{db: db}
}

func (up *UserPostgres) CreateUser(user shop.User) (string, error) {
	var id string
	query := fmt.Sprintf("INSERT INTO %s (firstname, lastname) VALUES ($1, $2) RETURNING id", getTableName())
	row := up.db.QueryRow(query, user.FirstName, user.LastName)
	if err := row.Scan(&id); err != nil {
		return "", err
	}
	return id, nil
}

func getTableName() string {
	return fmt.Sprintf("%s.%s", os.Getenv(dbSchemaKey), usersTable)
}