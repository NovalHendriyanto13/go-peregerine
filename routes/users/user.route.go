// Package users on routes
package users

import (
	"github.com/gofiber/fiber/v2"
	"peregerine/app/controllers/api"
	"peregerine/systems/types/interfaces"
)

// UserRoutes are collection of user routes url
func UserRoutes (app *fiber.App) {
	var controller interfaces.IndexControllerInterface = api.UserController{}
	user := app.Group("/users")

	user.Get("/", controller.Index)
}