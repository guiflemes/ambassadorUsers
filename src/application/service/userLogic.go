package service

import (
	"fmt"
	"users/src/application/port/in"
	"users/src/application/port/out"
	"users/src/utils"

	"github.com/pkg/errors"
)

type userLogic struct {
	userRepo      out.UserRepository
	storeService  in.UserStoreUseCase
	updateService in.UserUpdateUseCase
	deleteService in.DeleteUseCase
	getService    in.GetUserQuery
}

func NewUserLogic(userRepo out.UserRepository) in.UserService {
	return &userLogic{
		userRepo: userRepo,
	}

}

func (u *userLogic) GetAll() ([]*in.UserRespBody, error) {
	return u.getService.GetAll()
}

func (u *userLogic) GetById(id string) (*in.UserRespBody, error) {
	return u.getService.GetById(id)
}

func (u *userLogic) Update(userReq *in.UserUpdateReq) (*in.UserRespBody, error) {
	return u.updateService.Update(in.UserUpdateReqToDomain(userReq))

}

func (u *userLogic) Delete(id string) error {
	return u.deleteService.Delete(id)
}

func (u *userLogic) Store(user_req *in.UserReqBody) (*in.UserRespBody, error) {

	if exists, _, _ := u.getService.GetByEmail(user_req.Email); exists {
		return nil, errors.Wrap(utils.ErrUserAlredyExists, fmt.Sprintf("the given %s already exists", user_req.Email))
	}

	userDomain := in.ToUserDomain(user_req)

	return u.storeService.Store(userDomain)
}
