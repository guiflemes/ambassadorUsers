package service

import (
	"fmt"
	"users/src/application/port/in"
	"users/src/application/port/out"
	"users/src/utils"

	"github.com/pkg/errors"
)

type userService struct {
	storeService  in.UserStoreUseCase
	updateService in.UserUpdateUseCase
	deleteService in.DeleteUseCase
	getService    in.GetUserQuery
}

func NewUserService(userRepo out.UserRepository) in.UserUseCase {
	return &userService{
		storeService:  &userStoreService{userRepo: userRepo},
		updateService: &userUpdateService{userRepo: userRepo},
		deleteService: &userDeleteService{userRepo: userRepo},
		getService:    &userGetService{userRepo: userRepo},
	}

}

func (u *userService) GetAll() ([]*in.UserRespBody, error) {
	return u.getService.GetAll()
}

func (u *userService) GetById(id string) (*in.UserRespBody, error) {
	return u.getService.GetById(id)
}

func (u *userService) Update(userReq *in.UserUpdateReq) (*in.UserRespBody, error) {
	return u.updateService.Update(in.UserUpdateReqToDomain(userReq))
}

func (u *userService) Delete(id string) error {
	return u.deleteService.Delete(id)
}

func (u *userService) Store(user_req *in.UserReqBody) (*in.UserRespBody, error) {

	if exists, _, _ := u.getService.GetByEmail(user_req.Email); exists {
		return nil, errors.Wrap(utils.ErrUserAlredyExists, fmt.Sprintf("the given %s already exists", user_req.Email))
	}

	userDomain := in.ToUserDomain(user_req)

	return u.storeService.Store(userDomain)
}
