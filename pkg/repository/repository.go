package repository

import (
	"github.com/jmoiron/sqlx"
	shop "local/shop/backend"
)

type User interface {
	CreateUser(user shop.User) (string, error)
}

type Order interface {

}

type Category interface {

}

type Product interface {

}

type Report interface {

}

type Repository struct {
	User
	Order
	Category
	Product
	Report
}

func NewRepository(db *sqlx.DB) *Repository  {
	return &Repository{
		User: NewUserPostgres(db),
	}
}