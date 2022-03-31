package in

import (
	"users/src/domain"
)

type GetUserQuery interface {
	GetAll() (domain.UsersList, error)
	GetById(id string) (*domain.User, error)
	GetByEmail(email string) (bool, *domain.User, error)
}
