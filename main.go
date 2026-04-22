// Package main
package main

import (
	"fmt"
	"log"
	"github.com/gofiber/fiber/v2"
	// "github.com/tmc/langchaingo/llms/ollama"
	"peregerine/routes"
	"peregerine/configs"
	DI "peregerine/DI"
	CoreSecurity "peregerine/systems/middlewares"
)

func main() {
	appPort := configs.AppPort
	appName := configs.AppName
	fmt.Println(appName)

	c, err := DI.Build()
	if err != nil {
		panic(err)
	}
	defer c.Logger.Log.Sync()

	app := fiber.New(fiber.Config{
		// Limit request body to 10MB — prevents memory exhaustion from large payloads
		BodyLimit: 10 * 1024 * 1024,
		// Centralized JSON error handler for consistent API responses
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError
			if e, ok := err.(*fiber.Error); ok {
				code = e.Code
			}
			return ctx.Status(code).JSON(fiber.Map{
				"success": false,
				"code":    code,
				"message": err.Error(),
				"data":    fiber.Map{},
			})
		},
	})

	CoreSecurity.CoreSecurity(app)

	routes.RegisterRoute(app, c)

	log.Fatal(app.Listen(":" + appPort))
}