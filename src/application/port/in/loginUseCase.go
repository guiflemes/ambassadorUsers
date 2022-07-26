package in

import (
	"context"
)

type LoginUseCase interface {
	Authenticate(ctx context.Context, email string, password string) (bool, *UserRespBody, error)
}
