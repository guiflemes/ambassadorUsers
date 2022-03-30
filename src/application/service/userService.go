package service

import (
	"users/src/application/port/in"
	"users/src/application/port/out"
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
	var res []*in.UserRespBody

	users, err := u.getService.GetAll()

	if err != nil {
		return nil, err
	}

	for _, u := range users {
		user_res := in.NewUserRespBody(u)
		res = append(res, user_res)
	}

	return res, nil
}

func (u *userService) GetById(id string) (*in.UserRespBody, error) {
	userDomain, err := u.getService.GetById(id)

	if err != nil {
		return nil, err
	}
	return in.NewUserRespBody(userDomain), nil
}

func (u *userService) Update(userReq *in.UserUpdateReq) (*in.UserRespBody, error) {
	userDomain, err := u.getService.GetById(userReq.Id)

	if err != nil {
		return nil, err
	}

	userDomain.FirstName = userReq.FirstName
	userDomain.LastName = userReq.LastName
	userDomain.Email = userReq.Email

	userDomain, err = u.updateService.Update(userDomain)

	if err != nil {
		return nil, err
	}

	return in.NewUserRespBody(userDomain), nil
}

func (u *userService) Delete(id string) error {
	return u.deleteService.Delete(id)
}

func (u *userService) Store(user_req *in.UserReqBody) (*in.UserRespBody, error) {

	if exists, _, err := u.getService.GetByEmail(user_req.Email); exists {
		return nil, err
	}

	userDomain, err := u.storeService.Store(in.ToUserDomain(user_req))

	if err != nil {
		return nil, err
	}

	return in.NewUserRespBody(userDomain), nil
}
