package main

import (
	"fmt"
	"log"
	"github.com/gofiber/fiber/v2"
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