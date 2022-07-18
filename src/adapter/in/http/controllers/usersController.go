package controllers

import (
	"net/http"
	useCase "users/src/application/port/in"
	"users/src/application/service"
	"users/src/utils/container"

	"users/src/adapter/in/http/transport"

	"github.com/gofiber/fiber/v2"
)

type UserController struct {
	userService  useCase.UserUseCase
	errorHandler HandlerErrorUseCase
	encoder      transport.Encoder
}

func NewUserController(ctr *container.Container) *UserController {
	encode := &transport.BaseEncode{}

	return &UserController{
		userService:  service.NewUserService(ctr.Repositories.User),
		errorHandler: &ErrorHandler{encoder: nil},
		encoder:      encode,
	}
}

func (ctl *UserController) CreateUser(c *fiber.Ctx) {
	ctx := c.Context()
	userReq := &useCase.UserReqBody{}

	if err := c.BodyParser(userReq); err != nil {
		ctl.errorHandler.HandleError(c, err, http.StatusBadRequest)
		return
	}

	userResp, err := ctl.userService.Store(ctx, userReq)

	if err != nil {
		ctl.errorHandler.HandleError(c, err, http.StatusUnprocessableEntity)
		return
	}

	payload := ctl.encoder.Encode(userResp, nil, true)
	transport.Send(c, payload, http.StatusCreated)
}

func (ctl *UserController) GetUser(c *fiber.Ctx) {
	ctx := c.Context()
	userID := c.Params("id")

	userRep, err := ctl.userService.GetById(ctx, userID)

	if err != nil {
		ctl.errorHandler.HandleError(c, err, http.StatusUnprocessableEntity)
	}

	payload := ctl.encoder.Encode(userRep, nil, true)
	transport.Send(c, payload, http.StatusOK)
}
