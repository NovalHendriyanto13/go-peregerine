package upload

import (
	"github.com/gofiber/fiber/v2"
	"peregerine/app/controllers/api"
	"peregerine/systems/types/interfaces"
	AppMiddleware "peregerine/app/middlewares"
	DI "peregerine/DI"
)

func UploadRoutes (app *fiber.App, di *DI.DiHandler) {
	controller := api.UploadController{
		Redis: di.Redis.RedisRepo,
	}

	var baseUserController interfaces.IndexControllerInterface = &controller
	
	upload := app.Group("/upload", AppMiddleware.JwtProtected())

	upload.Get("/", baseUserController.Index)
	upload.Post("/create", baseUserController.Create)
	upload.Get("/detail", baseUserController.Detail)
}