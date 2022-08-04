package controllers

import (
	"net/http"
	useCase "users/src/application/port/in"
	"users/src/application/service"
	"users/src/utils/container"

	"users/src/adapter/in/http/transport"

	"github.com/gofiber/fiber/v2"
)

type JwtTokenController struct {
	svc          useCase.JwtTokenUseCase
	errorHandler HandlerErrorUseCase
	encoder      transport.Encoder
}

func NewJwtTokenController(ctr *container.Container) *JwtTokenController {
	encode := &transport.BaseEncode{}

	return &JwtTokenController{
		svc:          service.NewJwtTokenService(ctr.Repositories.User),
		errorHandler: &ErrorHandler{encoder: encode},
		encoder:      encode,
	}
}

// ShowUser godoc
// @Summary      Refresh token
// @Description  get auth token by refresh_token
// @Tags         login
// @Accept       json
// @Produce      json
// @Param        token body      in.JwtTokenRequest  true  "Token Body"
// @Success      200  {object}  transport.EncodedSuccess{data=auth.TokenPair,success=bool} "Result"
// @Failure      400  {string}  string    "Bad Request"
// @Failure      402  {string}  string    "Unauthorized"
// @Router       /api/v1/refresh_token [post]
func (t *JwtTokenController) RefreshToken(c *fiber.Ctx) error {
	ctx := c.Context()
	tokenReq := &useCase.JwtTokenRequest{}

	if err := c.BodyParser(tokenReq); err != nil {
		return t.errorHandler.HandleError(c, err, http.StatusBadRequest)
	}

	tokens, err := t.svc.RefreshToken(ctx, tokenReq)

	if err != nil {
		return t.errorHandler.HandleError(c, err, http.StatusUnauthorized)
	}
	payload := t.encoder.Encode(tokens, nil, true)
	return transport.Send(c, payload, http.StatusOK)
}
