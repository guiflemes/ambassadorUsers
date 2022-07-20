package transport

import (
	"github.com/gofiber/fiber/v2"
)

func Send(c *fiber.Ctx, payload *Encoded, code int) error {
	return c.Status(code).JSON(payload)
}
