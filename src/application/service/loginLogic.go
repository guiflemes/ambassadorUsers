package service

import (
	"users/src/application/port/in"
	"users/src/application/port/out"
)

type loginLogic struct {
	loginRepo out.LoginRepository
}

func NewLoginLogic(loginRepo out.LoginRepository) in.LoginService {
	return &loginLogic{
		loginRepo: loginRepo,
	}
}

func (l *loginLogic) Authenticate(email string, password string) (bool, *in.UserRespBody, error) {
	_, user, error := l.loginRepo.Authenticate(email, password)

	if error != nil {
		return false, nil, error
	}
	return true, in.NewUserRespBody(user), nil
}
