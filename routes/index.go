package routes

import (
	"github.com/gofiber/fiber/v2"
	"peregerine/routes/home"
	"peregerine/routes/users"
)

func RegisterRoute(app *fiber.App) {
	home.HomeRoutes(app)
	users.UserRoutes(app)
}