package startup

import (
	"log"
	"users/src/utils/config"
	"users/src/utils/container"

	"github.com/gofiber/fiber/v2"
)

func StartApp() {
	log.Print("StartApp")

	config := config.Parser()

	ctr, err := container.Resolve(config)

	if err != nil {
		panic(err)
	}

	app := fiber.New()

	setMiddleware(app)
	initRouters(app, ctr)
	app.Listen(":8000")

}
