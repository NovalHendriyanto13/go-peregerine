// Package main
package main

import (
	"fmt"
	"log"
	"github.com/gofiber/fiber/v2"
	// "github.com/tmc/langchaingo/llms/ollama"
	"peregerine/routes"
	"peregerine/configs"
)

func main() {
	appPort := configs.AppPort
	appName := configs.AppName
	fmt.Println(appName)
	app := fiber.New()

	routes.RegisterRoute(app)

	log.Fatal(app.Listen(":" + appPort))
}