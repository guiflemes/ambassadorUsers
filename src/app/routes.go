package app

import (
	"github.com/gofiber/fiber/v2"
)

func setup(app *fiber.App) {
	app.Group("api")
}
