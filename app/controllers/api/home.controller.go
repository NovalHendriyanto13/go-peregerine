// Package api main package to accommodate HomeController
package api

import (
	"github.com/gofiber/fiber/v2"
	"peregerine/systems/types/responses"
)

// HomeController is group of home routes action
type HomeController struct {
	responses.BaseResponse
}

// Index as an action from home routes to get list of data
func (h HomeController) Index(c *fiber.Ctx) error {
	resp := h.SuccessResponse(true, fiber.Map{
		"message": "index",
	})

	return c.JSON(resp)
}

// Create as an action from home routes to create data
func (h HomeController) Create(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "Index",
	})
}

// Detail as an action from home routes to specific data filtered by PK ID
func (h HomeController) Detail(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "Index",
	})
}

// Update as an action from home routes to Update data by PK ID
func (h HomeController) Update(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "Index",
	})
}

// Remove as an action from home routes to Delete data by PK ID
func (h HomeController) Remove(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "Index",
	})
}