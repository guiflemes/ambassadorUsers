package logic

import (
	"fmt"
	"users/src/domain"
	"users/src/dto"
	repo "users/src/repositories"
	srv "users/src/services"
	"users/src/utils"

	uuid "github.com/nu7hatch/gouuid"
	"github.com/pkg/errors"
)

type userLogic struct {
	userRepo repo.UserRepository
}

func NewUserLogic(userRepo repo.UserRepository) srv.UserService {
	return &userLogic{
		userRepo: userRepo,
	}

}

func (u *userLogic) GetAll() ([]*dto.UserRespBody, error) {
	var res []*dto.UserRespBody

	users, err := u.userRepo.GetAll()

	if err != nil {
		return nil, err
	}

	for _, u := range users {
		user_res := dto.NewUserRespBody(u)
		res = append(res, user_res)
	}

	return res, nil
}

func (u *userLogic) GetById(id string) (*dto.UserRespBody, error) {
	user, err := u.userRepo.GetBy(map[string]interface{}{"ID": id})
	if err != nil {
		return nil, err
	}

	res := dto.NewUserRespBody(user)

	return res, nil
}

func (u *userLogic) Store(user_req *dto.UserReqBody) (*dto.UserRespBody, error) {

	if exists, _, _ := u.getByEmail(user_req.Email); exists {
		return nil, errors.Wrap(utils.ErrUserAlredyExists, fmt.Sprintf("email %s already exists", user_req.Email))
	}

	user := user_req.ToUserDomain()

	if user.Id == "" {
		uid, _ := uuid.NewV4()
		user.Id = uid.String()
	}

	user.Password = repo.EncryptPassword(user.Password)

	user, err := u.userRepo.Store(user)

	if err != nil {
		return nil, err
	}

	return dto.NewUserRespBody(user), nil
}

func (u *userLogic) Update(user_req *dto.UserReqBody) (*dto.UserRespBody, error) {
	user, err := u.userRepo.Update(user_req.ToUserDomain())
	if err != nil {
		return nil, err
	}

	return dto.NewUserRespBody(user), nil
}

func (u *userLogic) Delete(id string) error {
	if id == "" {
		return errors.Wrap(utils.ErrInvalidParamater, "id can't be empty")
	}

	if err := u.userRepo.Delete(id); err != nil {
		return err
	}
	return nil
}

func (u *userLogic) getByEmail(email string) (bool, *domain.User, error) {
	res, err := u.userRepo.GetBy(map[string]interface{}{"Email": email})

	if err != nil {
		return false, res, err
	}

	return true, res, nil
}
