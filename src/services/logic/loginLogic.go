package logic

import (
	"users/src/domain"
	repo "users/src/repositories"
	svc "users/src/services"
)

type loginService struct {
	loginRepo repo.LoginRepository
}

func NewLoginLogic(loginRepo repo.LoginRepository) svc.LoginService {
	return &loginService{
		loginRepo,
	}
}

func (l *loginService) Authenticate(email string, password string) (bool, *domain.User, error) {
	return l.loginRepo.Authenticate(email, password)
}
