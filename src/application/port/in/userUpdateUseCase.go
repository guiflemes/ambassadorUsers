package in

import (
	"context"
	"users/src/domain"
)

type UserUpdateUseCase interface {
	Update(context.Context, *domain.User) (*domain.User, error)
}
