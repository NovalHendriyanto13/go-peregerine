package scrape

import (
	"github.com/gofiber/fiber/v2"
	"peregerine/app/controllers/api"
	"peregerine/systems/types/interfaces"
	DI "peregerine/DI"
)

func ScrapeRoute(app *fiber.App, di *DI.DiHandler) {
	controller := api.ScrapeController{
		Logger: di.Logger,
	}

	var baseScrapeController interfaces.IndexControllerInterface = &controller

	scrape := app.Group("/scrape")

	scrape.Post("/", baseScrapeController.Index)
}