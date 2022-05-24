package settings

import (
	"fmt"
	"users/src/adapter/out/persistence"

	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func StartApp() {
	log.Print("StartApp")

	sHost := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		GETENV("POSTGRES_HOST"),
		GETENV("POSTGRES_PORT"),
		GETENV("POSTGRES_USER"),
		GETENV("POSTGRES_PASSWORD"),
		GETENV("POSTGRES_DB_NAME"),
	)

	_ = persistence.NewPostgresRepository(sHost)

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	setup(app)

	app.Listen(":8000")
}
