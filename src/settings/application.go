package settings

import (
	"users/src/adapter/out/persistence"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func StartApp() {

	_, _ = persistence.NewMySQLRepository(GETSTRING("MYSQL_URL"))

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	setup(app)

	app.Listen(":8000")
}
