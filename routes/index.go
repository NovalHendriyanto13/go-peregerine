// Package routes are main routes that can register other routes
package routes

import (
	"github.com/gofiber/fiber/v2"
	"peregerine/routes/home"
	"peregerine/routes/users"
	"peregerine/routes/upload"
	DI "peregerine/DI"
	AppMiddleware "peregerine/app/middlewares"
)

// RegisterRoute is function to register all routes applications
func RegisterRoute(app *fiber.App, di *DI.DiHandler) {
	app.Group("/", AppMiddleware.AppKey())
	
	users.UserRoutes(app, di)
	
	home.HomeRoutes(app)
	upload.UploadRoutes(app, di)
	
}