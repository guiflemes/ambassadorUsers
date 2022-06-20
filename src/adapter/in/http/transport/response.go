package transport

import (
	"github.com/gofiber/fiber/v2"
)

func Send(c *fiber.Ctx, payload *Encoder, code int) {
	c.Status(code).JSON(payload)
}
