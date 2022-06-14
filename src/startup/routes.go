package startup

import (
	"users/src/utils/container"

	"github.com/gofiber/fiber/v2"
)

func initRouters(app *fiber.App, ctr *container.Container) {
	app.Group("api")
}
