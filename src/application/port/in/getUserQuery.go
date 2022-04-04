package in

import (
	"context"
	"users/src/domain"
)

type GetUserQuery interface {
	GetAll(ctx context.Context) (domain.UsersList, error)
	GetById(ctx context.Context, id string) (*domain.User, error)
	GetByEmail(ctx context.Context, email string) (bool, *domain.User, error)
}
