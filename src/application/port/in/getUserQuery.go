package in

import "users/src/domain"

type GetUserQuery interface {
	GetAll() ([]*UserRespBody, error)
	GetById(id string) (*UserRespBody, error)
	GetByEmail(email string) (bool, *domain.User, error)
}
