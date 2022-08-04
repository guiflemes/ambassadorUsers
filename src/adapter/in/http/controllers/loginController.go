package controllers

import (
	"net/http"
	useCase "users/src/application/port/in"
	"users/src/application/service"
	"users/src/utils"

	"users/src/utils/auth"
	"users/src/utils/container"

	"users/src/adapter/in/http/transport"

	"github.com/gofiber/fiber/v2"
)

type LoginController struct {
	loginSvc     useCase.LoginUseCase
	errorHandler HandlerErrorUseCase
	encoder      transport.Encoder
	jwtToken     auth.JwtToken
}

func NewLoginController(ctr *container.Container) *LoginController {
	encode := &transport.BaseEncode{}

	return &LoginController{
		loginSvc:     service.NewLoginService(ctr.Repositories.User, service.IsPasswordMatch),
		errorHandler: &ErrorHandler{encoder: encode},
		encoder:      encode,
		jwtToken:     auth.GenerateTokenPair,
	}
}

type userLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// ShowUser godoc
// @Summary      Login
// @Tags         login
// @Accept       json
// @Produce      json
// @Param        user body      userLogin  true  "Login"
// @Success      200  {object}  transport.EncodedSuccess{data=auth.TokenPair,success=bool} "Result"
// @Failure      422  {object}  transport.EncodedFail{error=string,success=bool} "UnprocessableEntity"
// @Failure      400  {string}  string    "Bad Request"
// @Failure      402  {string}  string    "Unauthorized"
// @Router       /api/v1/login [post]
func (ctl *LoginController) Login(c *fiber.Ctx) error {
	ctx := c.Context()

	userReq := &userLogin{}

	if err := c.BodyParser(userReq); err != nil {
		return ctl.errorHandler.HandleError(c, err, http.StatusBadRequest)
	}

	auth, userResp, err := ctl.loginSvc.Authenticate(ctx, userReq.Email, userReq.Password)

	if err != nil {
		return ctl.errorHandler.HandleError(c, err, http.StatusUnprocessableEntity)
	}

	if !auth {
		return ctl.errorHandler.HandleError(c, utils.ErrUnauthorized, http.StatusUnauthorized)
	}

	tokens, err := ctl.jwtToken(userResp)
	if err != nil {
		return ctl.errorHandler.HandleError(c, err, http.StatusBadRequest)
	}

	payload := ctl.encoder.Encode(tokens, nil, true)
	return transport.Send(c, payload, http.StatusOK)

}
