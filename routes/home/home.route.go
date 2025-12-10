package home

import (
	"github.com/gofiber/fiber/v2"
	"peregerine/app/controllers/api"
	"peregerine/systems/types/interfaces"
)

func HomeRoutes (app *fiber.App) {
	var controller interfaces.IndexControllerInterface = api.HomeController{}
	home := app.Group("/")

	home.Get("/", controller.Index)
}