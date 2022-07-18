package startup

import (
	"users/src/adapter/in/http/controllers"
	"users/src/utils/container"

	"github.com/gofiber/fiber/v2"
)

func initRouters(app *fiber.App, ctr *container.Container) {
	api := app.Group("api")

	api.Group("/v1")

	controllers.NewUserController(ctr)

}
