package in

import (
	"context"
	"users/src/domain"
)

type LoginUseCase interface {
	Authenticate(ctx context.Context, email string, password string) (bool, *domain.User, error)
}
