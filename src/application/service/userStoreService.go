package service

import (
	"users/src/application/port/in"
	"users/src/application/port/out"
	"users/src/domain"

	uuid "github.com/nu7hatch/gouuid"
	"github.com/pkg/errors"
)

type userStoreService struct {
	userRepo out.UserRepository
}

func (s *userStoreService) setId(userDomain *domain.User) {
	uid, _ := uuid.NewV4()
	userDomain.Id = uid.String()
}

func (s *userStoreService) encryptPassword(userDomain *domain.User) {
	userDomain.Password = EncryptPassword(userDomain.Password)
}

func (s *userStoreService) Store(userDomain *domain.User) (*in.UserRespBody, error) {

	s.setId(userDomain)
	s.encryptPassword(userDomain)

	if is_valid, err := userDomain.IsValid(); !is_valid {
		return nil, errors.Wrap(err, "user domais is not valid")
	}

	user, err := s.userRepo.Store(userDomain)

	if err != nil {
		return nil, errors.Wrap(err, "it was not possible to save User on Database")
	}

	return in.NewUserRespBody(user), nil

}
