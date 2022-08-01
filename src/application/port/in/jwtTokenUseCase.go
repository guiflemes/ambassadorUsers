package in

import (
	"context"
	"users/src/utils/auth"
)

type JwtTokenUseCase interface {
	RefreshToken(context.Context, *JwtTokenRequest) (*auth.TokenPair, error)
}
