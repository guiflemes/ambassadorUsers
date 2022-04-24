package service

import (
	"context"
	"users/src/application/port/out"
	"users/src/domain"

	"github.com/pkg/errors"
)

type userStoreService struct {
	userRepo out.UserRepository
}

func (s *userStoreService) encryptPassword(userDomain *domain.User) {
	userDomain.Password = EncryptPassword(userDomain.Password)
}

func (s *userStoreService) Store(ctx context.Context, userDomain *domain.User) (*domain.User, error) {

	s.encryptPassword(userDomain)

	if is_valid, err := userDomain.IsValid(); !is_valid {
		return nil, errors.Wrap(err, "user domais is not valid")
	}

	user, err := s.userRepo.Store(ctx, userDomain)

	if err != nil {
		return nil, errors.Wrap(err, "it was not possible to save User on Database")
	}

	return user, nil
}
