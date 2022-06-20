package http

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

type HandlerErrorUseCase interface {
	HandleError(*fiber.Ctx, error, int)
}

type ErrorHandler struct{}

func NewErrorHandler() *ErrorHandler {
	return &ErrorHandler{}
}

func (f *ErrorHandler) HandleError(c *fiber.Ctx, err error, code int) {
	log.Print(err)
	c.Status(code).SendString(err.Error())
}
