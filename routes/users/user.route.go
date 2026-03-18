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
		Logger: di.Logger,
	}
	var baseUserController interfaces.IndexControllerInterface = &controller
	var userController = &controller
	user := app.Group("/users")

	user.Get("/", baseUserController.Index)
	user.Post("/create", baseUserController.Create)
	user.Post("/login", userController.Login)
}