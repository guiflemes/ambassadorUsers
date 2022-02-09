package app

import (
	"users/src/repositories/mysql"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"

	"users/src/config"
)

func StartApp() {

	_, _ = mysql.NewMySQLRepository(config.GETSTRING("MYSQL_URL"))

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	setup(app)

	app.Listen(":8000")
}
