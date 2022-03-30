package in

import (
	"users/src/domain"
)

type UserUpdateUseCase interface {
	Update(*domain.User) (*domain.User, error)
}
