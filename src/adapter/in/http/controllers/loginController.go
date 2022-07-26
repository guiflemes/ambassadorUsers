package controllers

import (
	"net/http"
	useCase "users/src/application/port/in"
	"users/src/application/service"
	"users/src/utils"

	"users/src/utils/container"

	"users/src/adapter/in/http/transport"

	"github.com/gofiber/fiber/v2"
)

type LoginController struct {
	loginSvc     useCase.LoginUseCase
	errorHandler HandlerErrorUseCase
	encoder      transport.Encoder
}

func NewLoginController(ctr *container.Container) *LoginController {
	encode := &transport.BaseEncode{}

	return &LoginController{
		loginSvc:     service.NewLoginService(ctr.Repositories.User, service.IsPasswordMatch),
		errorHandler: &ErrorHandler{encoder: encode},
		encoder:      encode,
	}
}

func (ctl *LoginController) Login(c *fiber.Ctx) error {
	ctx := c.Context()
	email := c.FormValue("email")
	password := c.FormValue("password")

	if email == "" || password == "" {
		return ctl.errorHandler.HandleError(c, utils.ErrInvalidParameter, http.StatusBadRequest)
	}

	auth, userResp, err := ctl.loginSvc.Authenticate(ctx, email, password)

	if err != nil {
		return ctl.errorHandler.HandleError(c, err, http.StatusUnprocessableEntity)
	}

	if !auth {
		return ctl.errorHandler.HandleError(c, utils.ErrUnauthorized, http.StatusUnauthorized)
	}

	payload := ctl.encoder.Encode(userResp, nil, true)
	return transport.Send(c, payload, http.StatusAccepted)

}
