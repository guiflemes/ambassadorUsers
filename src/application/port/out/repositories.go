package out

import (
	"users/src/domain"
)

type UserRepository interface {
	GetAll() (domain.UsersList, error)
	GetBy(filter map[string]interface{}) (*domain.User, error)
	Store(data *domain.User) (*domain.User, error)
	Update(data *domain.User) (*domain.User, error)
	Delete(id string) error
}

type LoginRepository interface {
	Authenticate(username, password string) (bool, *domain.User, error)
}
