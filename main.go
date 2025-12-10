package main

import (
	"github.com/gofiber/fiber/v2"
	"peregerine/routes"
)

func main() {
	app := fiber.New()

	routes.RegisterRoute(app)

	app.Listen(":8000")
}