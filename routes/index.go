// Package routes are main routes that can register other routes
package routes

import (
	"github.com/gofiber/fiber/v2"
	"peregerine/routes/home"
	"peregerine/routes/users"
	DI "peregerine/DI"
)

// RegisterRoute is function to register all routes applications
func RegisterRoute(app *fiber.App, di *DI.DiHandler) {
	home.HomeRoutes(app)
	users.UserRoutes(app, di)
}