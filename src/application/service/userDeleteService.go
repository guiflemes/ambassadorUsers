package service

import (
	"context"
	"users/src/application/port/out"
	"users/src/utils"

	"github.com/pkg/errors"
)

type userDeleteService struct {
	userRepo out.UserRepository
}

func (s *userDeleteService) Delete(ctx context.Context, id string) error {
	if id == "" {
		return errors.Wrap(utils.ErrInvalidParameter, "id can't be empty")
	}

	if err := s.userRepo.Delete(ctx, id); err != nil {
		return errors.Wrap(err, "it was not possible to delete user")
	}

	return nil

}
