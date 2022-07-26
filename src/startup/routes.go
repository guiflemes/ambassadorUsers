package startup

import (
	"users/src/adapter/in/http/controllers"
	"users/src/utils/container"

	"github.com/gofiber/fiber/v2"
)

func initRouters(app *fiber.App, ctr *container.Container) {
	apiV1 := app.Group("api/v1")

	loginClt := controllers.NewLoginController(ctr)
	userClt := controllers.NewUserController(ctr)

	apiV1.Post("/login", loginClt.Login)

	apiV1.Get("/users/:id", userClt.GetUser)
	apiV1.Post("/users", userClt.CreateUser)
	apiV1.Put("/users/:id", userClt.UpdateUser)

}
