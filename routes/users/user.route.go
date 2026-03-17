// Package users on routes
package users

import (
	"github.com/gofiber/fiber/v2"
	"peregerine/app/controllers/api"
	"peregerine/systems/types/interfaces"
	DI "peregerine/DI"
)

// UserRoutes are collection of user routes url
func UserRoutes (app *fiber.App, di *DI.DiHandler) {
	controller := api.UserController{
		Redis: di.Redis.RedisRepo,
	}
	var userController interfaces.IndexControllerInterface = &controller
	user := app.Group("/users")

	user.Get("/", userController.Index)
	user.Post("/create", userController.Create)
}