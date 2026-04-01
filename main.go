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

	app := fiber.New()

	CoreSecurity.CoreSecurity(app)

	routes.RegisterRoute(app, c)

	log.Fatal(app.Listen(":" + appPort))
}