package web

import (
	useCase "users/src/application/port/in"
	svc "users/src/application/service"
)

type UserController struct {
	usecase useCase.UserUseCase
}

func NewUserController() UserController {
	ctl := UserController{
		usecase : svc.NewUserService()
	}
}