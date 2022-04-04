package in

import (
	"context"
)

type UserUseCase interface {
	GetAll(ctx context.Context) ([]*UserRespBody, error)
	GetById(ctx context.Context, id string) (*UserRespBody, error)
	Store(ctx context.Context, u *UserReqBody) (*UserRespBody, error)
	Update(ctx context.Context, u *UserUpdateReq) (*UserRespBody, error)
	Delete(ctx context.Context, id string) error
}
