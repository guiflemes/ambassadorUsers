package web

import (
	useCase "users/src/application/port/in"
)

type UserController struct {
	usecase useCase.UserUseCase
}
