package logic

import (
	"fmt"
	"users/src/domain"
	repo "users/src/repositories"
	srv "users/src/services"
	"users/src/utils"

	uuid "github.com/nu7hatch/gouuid"
	"github.com/pkg/errors"
	"gopkg.in/dealancer/validate.v2"
)

type userLogic struct {
	userRepo repo.UserRepository
}

func NewUserLogic(userRepo repo.UserRepository) srv.UserService {
	return &userLogic{
		userRepo: userRepo,
	}
}

func (u *userLogic) GetAll() ([]domain.User, error) {
	res := []domain.User{}
	if res, err := u.userRepo.GetAll(); err != nil {
		return res, err
	}
	return res, nil
}

func (u *userLogic) GetById(id string) (*domain.User, error) {
	res, err := u.userRepo.GetBy(map[string]interface{}{"ID": id})
	if err != nil {
		return res, err
	}
	return res, nil
}

func (u *userLogic) Store(data *domain.User) error {
	if err := validate.Validate(data); err != nil {
		return errors.Wrap(utils.ErrUserInvalid, "error trying to create user")
	}

	if data.Id == "" {
		uid, _ := uuid.NewV4()
		data.Id = uid.String()
	}

	if exists, _, _ := u.GetByEmail(data.Email); exists {
		return errors.Wrap(utils.ErrUserAlredyExists, fmt.Sprintf("email %s already exists", data.Email))
	}

	data.Password = repo.EncryptPassword(data.Password)

	return u.userRepo.Store(data)
}

func (u *userLogic) Update(id string, data map[string]interface{}) (*domain.User, error) {
	return u.userRepo.Update(id, data)
}

func (u *userLogic) Delete(id string) error {
	return u.userRepo.Delete(id)
}

func (u *userLogic) GetByEmail(email string) (bool, *domain.User, error) {
	res, err := u.userRepo.GetBy(map[string]interface{}{"Email": email})

	if err != nil {
		return false, res, err
	}

	return true, res, nil
}
