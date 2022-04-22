package service

import (
	shop "local/shop/backend"
	"local/shop/backend/pkg/repository"
)

type UserService struct {
	repo repository.User
}

func NewUserService(repo repository.User) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) CreateUser(user shop.User) (string, error) {
	return s.repo.CreateUser(user)
}
