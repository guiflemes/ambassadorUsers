package service

import (
	"users/src/application/port/in"
	"users/src/application/port/out"
	"users/src/domain"
)

type userUpdateService struct {
	userRepo out.UserRepository
}

func (s *userUpdateService) Update(userDomain *domain.User) (*in.UserRespBody, error) {
	//TODO validate domain and change how it updates
	user, err := s.userRepo.Update(userDomain)
	if err != nil {
		return nil, err
	}

	return in.NewUserRespBody(user), nil

}
