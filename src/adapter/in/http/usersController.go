package http

import (
	"net/http"
	useCase "users/src/application/port/in"
	"users/src/application/service"
	"users/src/utils/container"

	"users/src/adapter/in/http/transport"

	"github.com/gofiber/fiber/v2"
)

type UserController struct {
	usecase      useCase.UserUseCase
	errorHandler HandlerErrorUseCase
	encoder      transport.Encoder
}

func NewUserController(ctr *container.Container) *UserController {
	encode := &transport.BaseEncode{}

	return &UserController{
		usecase:      service.NewUserService(ctr.Repositories.User),
		errorHandler: &ErrorHandler{encoder: encode},
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

	userResp, err := ctl.usecase.Store(ctx, userReq)

	if err != nil {
		ctl.errorHandler.HandleError(c, err, http.StatusUnprocessableEntity)
		return
	}

	payload := ctl.encoder.Encode(userResp, nil, "true")
	transport.Send(c, payload, http.StatusOK)
}
