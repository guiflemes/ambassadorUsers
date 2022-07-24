package service

import (
	"users/src/application/port/out"
	"users/src/domain"

	"context"

	"github.com/pkg/errors"
)

type userUpdateService struct {
	userRepo out.UserRepository
}

func (s *userUpdateService) Update(ctx context.Context, userDomain *domain.User) (*domain.User, error) {

	if is_valid, err := userDomain.IsValid(); !is_valid {
		return nil, errors.Wrap(err, "user domain is not valid")
	}

	user, err := s.userRepo.Update(ctx, userDomain)

	if err != nil {
		return nil, err
	}

	return user, nil

}
