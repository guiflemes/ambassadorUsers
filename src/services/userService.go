package services

import (
	"users/src/domain"
)

type UserService interface {
	GetAll() ([]domain.User, error)
	GetById(id string) (*domain.User, error)
	Store(data *domain.User) error
	Update(*domain.User) (*domain.User, error)
	Delete(id string) error
}
