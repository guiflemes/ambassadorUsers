package startup

import (
	"users/src/adapter/in/http/controllers"
	"users/src/utils/container"

	"github.com/gofiber/fiber/v2"
)

func initRouters(app *fiber.App, ctr *container.Container) {
	apiV1 := app.Group("api/v1")

	loginCtl := controllers.NewLoginController(ctr)
	userCtl := controllers.NewUserController(ctr)
	jwtTokenCtl := controllers.NewJwtTokenController(ctr)

	apiV1.Post("/login", loginCtl.Login)
	apiV1.Post("/users", userCtl.CreateUser)
	apiV1.Post("/refresh_token", jwtTokenCtl.RefreshToken)

	apiV1.Use(isLoggedIn)

	apiV1.Get("/users/:id", loggedInUser, userCtl.GetUser)
	apiV1.Put("/users/:id", loggedInUser, userCtl.UpdateUser)

}
