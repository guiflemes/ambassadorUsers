package repositories

import (
	"users/src/domain"
)

type UserRepository interface {
	GetAll() ([]domain.User, error)
	GetBy(filter map[string]interface{}) (*domain.User, error)
	Store(data *domain.User) error
	Update(data map[string]interface{}, id string) (*domain.User, error)
	Delete(id string) error
	Authenticate(username, password string) (bool, *domain.User, error)
}
