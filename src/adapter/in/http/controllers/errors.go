package controllers

import (
	"users/src/adapter/in/http/transport"

	"github.com/gofiber/fiber/v2"
)

type HandlerErrorUseCase interface {
	HandleError(*fiber.Ctx, error, int) error
}

type ErrorHandler struct {
	encoder transport.Encoder
}

func (e *ErrorHandler) HandleError(c *fiber.Ctx, err error, code int) error {
	payload := e.encoder.Encode(nil, err.Error(), false)
	return transport.Send(c, payload, code)
}
