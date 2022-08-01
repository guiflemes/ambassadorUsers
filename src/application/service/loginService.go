package service

import (
	"context"
	"fmt"
	"users/src/application/port/out"
	"users/src/domain"

	"users/src/utils"

	"strings"

	"github.com/pkg/errors"
)

type passwordMatch func(password, userpass string) bool

type LoginService struct {
	userRepo out.GetterBy
	passMath passwordMatch
}

func NewLoginService(userRepo out.GetterBy, passMath passwordMatch) *LoginService {
	return &LoginService{
		userRepo: userRepo,
		passMath: passMath,
	}
}

func (l *LoginService) Authenticate(ctx context.Context, email string, password string) (bool, *domain.User, error) {

	if strings.TrimSpace(email) == "" || strings.TrimSpace(password) == "" {
		return false, nil, utils.ErrInvalidParameter
	}

	user, err := l.userRepo.GetBy(ctx, map[string]interface{}{"email": email})

	if err != nil {
		return false, nil, errors.New(fmt.Sprintf("the given email %s doest not exists", email))
	}

	if !l.passMath(password, user.Password) {
		return false, nil, errors.New("Password does not match")
	}

	return true, user, nil
}
