package service

import (
	"fmt"
	"users/src/application/port/out"
	"users/src/domain"
	"users/src/utils"

	"github.com/pkg/errors"
)

type userGetService struct {
	userRepo out.UserRepository
}

func (s *userGetService) GetAll() ([]*domain.User, error) {
	users, err := s.userRepo.GetAll()

	if err != nil {
		return nil, errors.Wrap(err, "error retrieving all users")
	}

	return users, err
}

func (s *userGetService) GetById(id string) (*domain.User, error) {
	user, err := s.userRepo.GetBy(map[string]interface{}{"ID": id})
	if err != nil {
		return nil, errors.Wrap(utils.ErrUserNotFound, fmt.Sprintf("the given %s doest not exists", id))
	}
	return user, nil
}

func (s *userGetService) GetByEmail(email string) (bool, *domain.User, error) {
	user, err := s.userRepo.GetBy(map[string]interface{}{"Email": email})

	if err != nil {
		return false, user, errors.Wrap(utils.ErrUserNotFound, fmt.Sprintf("the given %s doest not exists", email))
	}

	return true, nil, nil
}
