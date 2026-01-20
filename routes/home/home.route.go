// Package Home as routes
package home

import (
	"github.com/gofiber/fiber/v2"
	"peregerine/app/controllers/api"
	"peregerine/systems/types/interfaces"
)

// HomeRoutes are the list of home routes url
func HomeRoutes (app *fiber.App) {
	var controller interfaces.IndexControllerInterface = api.HomeController{}
	home := app.Group("/")

	home.Get("/", controller.Index)
	home.Post("/create", controller.Create)
}