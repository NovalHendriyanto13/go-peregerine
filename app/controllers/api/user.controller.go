// Package api main package to accommodate UserController
package api

import (
	"github.com/gofiber/fiber/v2"
	"peregerine/systems/types/responses"
	RedisService "peregerine/systems/services/redis/types"
)

// UserController is group of home routes action
type UserController struct {
	responses.BaseResponse
	Redis RedisService.SettingService
}

// Index as an action from user routes to get list of data
func (u UserController) Index(c *fiber.Ctx) error {
	err := u.Redis.Set("users", "user", "test")
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	resp := u.SuccessResponse(true, fiber.Map{
		"message": "user dong",
	})

	return c.JSON(resp)
}

// Create as an action from User routes to get list of data
func (u UserController) Create(c *fiber.Ctx) error {
	val, _ := u.Redis.Get("users", "user")

	return c.JSON(fiber.Map{
		"message": val,
	})
}

// Detail as an action from User routes to get list of data
func (u UserController) Detail(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "Index",
	})
}

// Update as an action from User routes to get list of data
func (u UserController) Update(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "Index",
	})
}

// Remove as an action from User routes to get list of data
func (u UserController) Remove(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "Index",
	})
}