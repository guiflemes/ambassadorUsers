package in

import (
	"context"
	"users/src/domain"
)

type UserStoreUseCase interface {
	Store(context.Context, *domain.User) (*domain.User, error)
}
