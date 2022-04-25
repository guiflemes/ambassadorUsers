package service

import (
	"fmt"
	"users/src/application/port/out"
	"users/src/domain"
	"users/src/utils"

	"context"

	"github.com/pkg/errors"
)

type userGetService struct {
	userRepo out.UserRepository
}

func (s *userGetService) GetAll(ctx context.Context) (domain.UsersList, error) {
	users, err := s.userRepo.GetAll(ctx)

	if err != nil {
		return nil, errors.Wrap(err, "error retrieving all users")
	}

	return users, err
}

func (s *userGetService) GetById(ctx context.Context, id string) (*domain.User, error) {
	user, err := s.userRepo.GetBy(ctx, map[string]interface{}{"id": id})
	if err != nil {
		return nil, errors.Wrap(utils.ErrUserNotFound, fmt.Sprintf("the given %s doest not exists", id))
	}
	return user, nil
}

func (s *userGetService) GetByEmail(ctx context.Context, email string) (bool, *domain.User, error) {
	user, err := s.userRepo.GetBy(ctx, map[string]interface{}{"email": email})

	if err != nil {
		return false, user, errors.Wrap(utils.ErrUserNotFound, fmt.Sprintf("the given %s doest not exists", email))
	}

	return true, nil, nil
}
