package startup

import (
	"github.com/gofiber/fiber/v2"

	"users/src/utils"

	"github.com/gofiber/fiber/v2/middleware/cors"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/golang-jwt/jwt/v4"
)

var isLoggedIn = jwtware.New(jwtware.Config{
	SigningKey: []byte("secret"),
})

func loggedInUser(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	if c.Params("id") != claims["id"] {
		return utils.ErrUnauthorized
	}

	return c.Next()

}

func setMiddleware(app *fiber.App) {
	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))
}
