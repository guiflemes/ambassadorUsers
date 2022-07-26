package out

import (
	"context"
	"users/src/domain"
)

type UserRepository interface {
	GetAll(context.Context) (domain.UsersList, error)
	GetterBy
	Store(context.Context, *domain.User) (*domain.User, error)
	Update(context.Context, *domain.User) (*domain.User, error)
	Delete(ctx context.Context, id string) error
}

type LoginRepository interface {
	Authenticate(ctx context.Context, username, password string) (bool, *domain.User, error)
}

type GetterBy interface {
	GetBy(ctx context.Context, filter map[string]interface{}) (*domain.User, error)
}
