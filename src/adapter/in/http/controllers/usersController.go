package controllers

import (
	"net/http"
	useCase "users/src/application/port/in"
	"users/src/application/service"
	"users/src/utils/container"

	"users/src/adapter/in/http/transport"

	"github.com/gofiber/fiber/v2"
)

type UserController interface {
	CreateUser(c *fiber.Ctx) error
	GetUser(c *fiber.Ctx) error
}

type UserControllerDefault struct {
	userService  useCase.UserUseCase
	errorHandler HandlerErrorUseCase
	encoder      transport.Encoder
}

func NewUserController(ctr *container.Container) *UserControllerDefault {
	encode := &transport.BaseEncode{}

	return &UserControllerDefault{
		userService:  service.NewUserService(ctr.Repositories.User),
		errorHandler: &ErrorHandler{encoder: encode},
		encoder:      encode,
	}
}

func (ctl *UserControllerDefault) CreateUser(c *fiber.Ctx) error {
	ctx := c.Context()
	userReq := &useCase.UserReqBody{}

	if err := c.BodyParser(userReq); err != nil {
		return ctl.errorHandler.HandleError(c, err, http.StatusBadRequest)
	}

	userResp, err := ctl.userService.Store(ctx, userReq)

	if err != nil {
		return ctl.errorHandler.HandleError(c, err, http.StatusUnprocessableEntity)
	}

	payload := ctl.encoder.Encode(userResp, nil, true)
	return transport.Send(c, payload, http.StatusCreated)
}

func (ctl *UserControllerDefault) GetUser(c *fiber.Ctx) error {
	ctx := c.Context()
	userID := c.Params("id")

	userRep, err := ctl.userService.GetById(ctx, userID)

	if err != nil {
		return ctl.errorHandler.HandleError(c, err, http.StatusUnprocessableEntity)
	}

	payload := ctl.encoder.Encode(userRep, nil, true)
	return transport.Send(c, payload, http.StatusOK)
}

func (ctl *UserControllerDefault) UpdateUser(c *fiber.Ctx) error {
	ctx := c.Context()
	userID := c.Params("id")

	userUpdateReq := &useCase.UserUpdateReq{}
	userUpdateReq.Id = userID

	if err := c.BodyParser(userUpdateReq); err != nil {
		return ctl.errorHandler.HandleError(c, err, http.StatusBadRequest)
	}

	userResp, err := ctl.userService.Update(ctx, userUpdateReq)

	if err != nil {
		return ctl.errorHandler.HandleError(c, err, http.StatusUnprocessableEntity)
	}

	payload := ctl.encoder.Encode(userResp, nil, true)
	return transport.Send(c, payload, http.StatusCreated)

}
