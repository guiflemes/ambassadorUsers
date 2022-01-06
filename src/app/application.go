package app

import (
	"users/src/database"

	"github.com/gofiber/fiber/v2"
)

func StartApp() {
	database.StartDB()

	app := fiber.New()

	app.Listen(":8000")
}
