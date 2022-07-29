package auth

import (
	"time"
	"users/src/application/port/in"

	"github.com/golang-jwt/jwt/v4"
)

type JwtToken func(user *in.UserRespBody) (*TokenPair, error)

func newAccessToken(user *in.UserRespBody) *jwt.Token {
	claims := jwt.MapClaims{
		"id":        user.Id,
		"firstName": user.FirstName,
		"lastName":  user.LastName,
		"email":     user.Email,
		"exp":       time.Now().Add(time.Minute * 15).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token
}

func newRefreshToken(user *in.UserRespBody) *jwt.Token {
	claims := jwt.MapClaims{
		"id":  user.Id,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token
}

type TokenPair struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func GenerateTokenPair(user *in.UserRespBody) (*TokenPair, error) {
	accessToken := newAccessToken(user)
	refreshToken := newRefreshToken(user)

	var asString []string

	tokenPair := TokenPair{}

	for _, token := range []*jwt.Token{
		accessToken,
		refreshToken,
	} {

		t, err := token.SignedString([]byte("secret"))

		if err != nil {
			return nil, err
		}

		asString = append(asString, t)
	}

	tokenPair.AccessToken = asString[0]
	tokenPair.RefreshToken = asString[1]
	return &tokenPair, nil
}
