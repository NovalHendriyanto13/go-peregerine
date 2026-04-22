package coremiddleware

import (
	"time"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

func CoreSecurity (app *fiber.App) {
	// FIX: Recover from panics — always first
	app.Use(recover.New(recover.Config{
		EnableStackTrace: true,
	}))

	// FIX: Add request ID to every request for tracing
	app.Use(requestid.New())

	app.Use(logger.New(logger.Config{
		Format: "[${time}] ${status} - ${latency} ${method} ${path} id=${locals:requestid}\n",
	}))

	app.Use(helmet.New())

	app.Use(cors.New(cors.Config{
		AllowMethods: "GET,POST,PUT,DELETE",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
	}))

	app.Use(limiter.New(limiter.Config{
		Max: 100,
		Expiration: 1 * time.Minute,
		KeyGenerator: func(c *fiber.Ctx) string {
			// Rate-limit per IP address
			return c.IP()
		},
		LimitReached: func(c *fiber.Ctx) error {
			return c.Status(fiber.StatusTooManyRequests).JSON(fiber.Map{
				"success": false,
				"code":    fiber.StatusTooManyRequests,
				"message": "Too many requests, please slow down",
				"data":    fiber.Map{},
			})
		},
	}))

	// FIX: Add gzip/deflate response compression
	app.Use(compress.New(compress.Config{
		Level: compress.LevelBestSpeed,
	}))
}