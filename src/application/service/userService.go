package service

import (
	"context"
	"users/src/application/port/in"
	"users/src/application/port/out"
	"users/src/domain"
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

func (u *userService) GetAll(ctx context.Context) ([]*in.UserRespBody, error) {
	var res []*in.UserRespBody

	users, err := u.getService.GetAll(ctx)

	if err != nil {
		return nil, err
	}

	for _, u := range users {
		user_res := in.NewUserRespBody(u)
		res = append(res, user_res)
	}

	return res, nil
}

func (u *userService) GetById(ctx context.Context, id string) (*in.UserRespBody, error) {
	userDomain, err := u.getService.GetById(ctx, id)

	if err != nil {
		return nil, err
	}
	return in.NewUserRespBody(userDomain), nil
}

func (u *userService) Update(ctx context.Context, userReq *in.UserUpdateReq) (*in.UserRespBody, error) {
	userDomain, err := u.getService.GetById(ctx, userReq.Id)

	if err != nil {
		return nil, err
	}

	userDomain.FirstName = userReq.FirstName
	userDomain.LastName = userReq.LastName
	userDomain.Email = userReq.Email

	userDomain, err = u.updateService.Update(ctx, userDomain)

	if err != nil {
		return nil, err
	}

	return in.NewUserRespBody(userDomain), nil
}

func (u *userService) Delete(ctx context.Context, id string) error {
	return u.deleteService.Delete(ctx, id)
}

func (u *userService) Store(ctx context.Context, user_req *in.UserReqBody) (*in.UserRespBody, error) {

	if exists, _, _ := u.getService.GetByEmail(ctx, user_req.Email); exists {
		return nil, errors.Wrap(utils.ErrUserAlreadyExists, user_req.Email)
	}

	d := in.ToUserDomain(user_req)
	d.IsActive = true
	d.Role = domain.Admin

	userDomain, err := u.storeService.Store(ctx, d)

	if err != nil {
		return nil, err
	}

	return in.NewUserRespBody(userDomain), nil
}
