package http

import (
	"net/http"
	useCase "users/src/application/port/in"
	"users/src/application/service"
	"users/src/utils/container"

	"github.com/gofiber/fiber/v2"
)

type UserController struct {
	usecase      useCase.UserUseCase
	errorHandler HandlerErrorUseCase
}

func NewUserController(ctr *container.Container) *UserController {
	return &UserController{
		usecase:      service.NewUserService(ctr.Repositories.User),
		errorHandler: NewErrorHandler(),
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

	c.Status(http.StatusOK).JSON(userResp)
}
