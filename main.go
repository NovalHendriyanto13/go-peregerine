// Package main
package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
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
	defer c.Close()

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

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	serverErr := make(chan error, 1)
	go func() {
		serverErr <- app.Listen(":" + appPort)
	}()

	select {
	case err := <-serverErr:
		log.Println(err)
	case <-ctx.Done():
		log.Println("Shutting down server...")
		// Stop accepting new requests and wait for in-flight ones before closing Redis/queue
		if err := app.ShutdownWithTimeout(10 * time.Second); err != nil {
			log.Println(err)
		}
	}
}