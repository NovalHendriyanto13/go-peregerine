package api

import (
	"github.com/gofiber/fiber/v2"
	"peregerine/systems/types/responses"
)

type UserController struct {
	responses.BaseResponse
}

func (u UserController) Index(c *fiber.Ctx) error {
	resp := u.SuccessResponse(true, fiber.Map{
		"message": "user dong",
	})

	return c.JSON(resp)
}

func (u UserController) Create(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "Index",
	})
}

func (u UserController) Detail(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "Index",
	})
}

func (u UserController) Update(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "Index",
	})
}

func (u UserController) Remove(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "Index",
	})
}