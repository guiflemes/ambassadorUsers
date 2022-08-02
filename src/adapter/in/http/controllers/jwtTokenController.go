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

func (t *JwtTokenController) RefreshToken(c *fiber.Ctx) error {
	ctx := c.Context()
	tokenReq := &useCase.JwtTokenRequest{}

	if err := c.BodyParser(tokenReq); err != nil {
		return t.errorHandler.HandleError(c, err, http.StatusBadGateway)
	}

	tokens, err := t.svc.RefreshToken(ctx, tokenReq)

	if err != nil {
		return t.errorHandler.HandleError(c, err, http.StatusUnauthorized)
	}
	payload := t.encoder.Encode(tokens, nil, true)
	return transport.Send(c, payload, http.StatusAccepted)
}
