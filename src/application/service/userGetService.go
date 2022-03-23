package service

import (
	"users/src/application/port/in"
	"users/src/application/port/out"
	"users/src/domain"
)

type getUserService struct {
	userRepo out.UserRepository
}

func (s *getUserService) GetAll() ([]*in.UserRespBody, error) {
	var res []*in.UserRespBody

	users, err := s.userRepo.GetAll()

	if err != nil {
		return nil, err
	}

	for _, u := range users {
		user_res := in.NewUserRespBody(u)
		res = append(res, user_res)
	}

	return res, nil
}

func (s *getUserService) GetById(id string) (*in.UserRespBody, error) {
	user, err := s.userRepo.GetBy(map[string]interface{}{"ID": id})
	if err != nil {
		return nil, err
	}

	res := in.NewUserRespBody(user)

	return res, nil
}

func (s *getUserService) GetByEmail(email string) (bool, *domain.User, error) {
	res, err := s.userRepo.GetBy(map[string]interface{}{"Email": email})

	if err != nil {
		return false, res, err
	}

	return true, res, nil
}
