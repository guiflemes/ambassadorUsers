package service

import (
	"context"
	"users/src/application/port/in"
	"users/src/utils"
	"users/src/utils/auth"

	"users/src/application/port/out"

	"github.com/golang-jwt/jwt/v4"
	"github.com/pkg/errors"
)

type TokenService struct {
	userRepo out.GetterBy
}

func (t *TokenService) RefreshToken(ctx context.Context, tokenReq *in.JwtTokenRequest) (*auth.TokenPair, error) {

	token, err := jwt.Parse(tokenReq.RefreshToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte("secret"), nil

	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {

		userID := claims["id"].(string)
		user, err := t.userRepo.GetBy(ctx, map[string]interface{}{"id": userID})

		if err != nil {
			return nil, err
		}

		newTokenPair, err := auth.GenerateTokenPair(user)

		if err != nil {
			return nil, utils.ErrUnauthorized
		}

		return newTokenPair, nil
	}

	return nil, err
}
