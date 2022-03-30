package in

import (
	"users/src/domain"
)

type UserStoreUseCase interface {
	Store(*domain.User) (*domain.User, error)
}
