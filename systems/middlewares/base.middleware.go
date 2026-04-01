package coremiddleware

import (
	"time"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/timeout"
)

func CoreSecurity (app *fiber.App) {
	app.Use(recover.New())

	app.Use(logger.New())

	app.Use(helmet.New())

	app.Use(cors.New(cors.Config{
		AllowMethods: "GET,POST,PUT,DELETE",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
	}))

	app.Use(limiter.New(limiter.Config{
		Max: 100,
		Expiration: 1 * time.Minute,
	}))

	// Timeout protection
	app.Use(timeout.NewWithContext(func(c *fiber.Ctx) error {
		return c.Next()
	}, 5*time.Second))
}