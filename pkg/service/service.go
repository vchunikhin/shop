package service

import "local/shop/backend/pkg/repository"

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

type Service struct {
	User
	Order
	Category
	Product
	Report
}

func NewService(rep *repository.Repository) *Service  {
	return &Service{}
}