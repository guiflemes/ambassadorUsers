package services

import "users/src/domain"

type UserService interface {
	GetAll() ([]domain.User, error)
	GetById(id string) (*domain.User, error)
	Store(data *domain.User) error
	Update(id string, data map[string]interface{}) (*domain.User, error)
	Delete(id string) error
}
