package repository

import "github.com/jmoiron/sqlx"

type User interface {

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
	return &Repository{}
}