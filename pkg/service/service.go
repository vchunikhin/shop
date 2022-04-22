package service

import (
	shop "local/shop/backend"
	"local/shop/backend/pkg/repository"
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

type Service struct {
	User
	Order
	Category
	Product
	Report
}

func NewService(rep *repository.Repository) *Service  {
	return &Service{
		User: NewUserService(rep),
	}
}