package startup

import (
	"users/src/adapter/in/http/controllers"
	"users/src/utils/container"

	"github.com/gofiber/fiber/v2"
)

func initRouters(app *fiber.App, ctr *container.Container) {
	apiV1 := app.Group("api/v1")

	clt := controllers.NewUserController(ctr)

	apiV1.Get("/users/:id", clt.GetUser)
	apiV1.Post("/users", clt.CreateUser)
	apiV1.Put("/users/:id", clt.UpdateUser)

}
