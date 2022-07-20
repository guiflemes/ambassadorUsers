package startup

import (
	"github.com/gofiber/fiber/v2"

	"github.com/gofiber/fiber/v2/middleware/cors"
)

func setMiddleware(app *fiber.App) {
	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))
	// app.Use(logger.New(logger.Config{
	// 	Next:         nil,
	// 	Format:       "[${time}] ${status} - ${latency} ${method} ${path}\n",
	// 	TimeFormat:   "15:04:05",
	// 	TimeZone:     "Local",
	// 	TimeInterval: 500 * time.Millisecond,
	// 	Output:       os.Stderr,
	// }))
	// app.Use(csrf.New())
	// app.Use(requestid.New())
}
