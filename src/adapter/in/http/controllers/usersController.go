package controllers

import (
	"net/http"
	useCase "users/src/application/port/in"
	"users/src/application/service"
	"users/src/utils/container"

	"users/src/adapter/in/http/transport"

	_ "users/docs"

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

// ShowUser godoc
// @Summary      Show an user
// @Description  get user by ID
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "User ID"
// @Security Authorization
// @in header
// @Success      200  {object}  transport.EncodedSuccess{data=in.UserRespBody,success=bool} "Result"
// @Failure      422  {object}  transport.EncodedFail{error=string,success=bool} "UnprocessableEntity"
// @Failure      400  {string}  string    "Bad Request"
// @Failure      402  {string}  string    "Unauthorized"
// @Router       /api/v1/users/{id} [get]
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
