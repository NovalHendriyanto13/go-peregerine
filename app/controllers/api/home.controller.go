package api

import (
	"github.com/gofiber/fiber/v2"
	"peregerine/systems/types/responses"
)

type HomeController struct {
	responses.BaseResponse
}

func (h HomeController) Index(c *fiber.Ctx) error {
	resp := h.SuccessResponse(true, fiber.Map{
		"message": "index",
	})

	return c.JSON(resp)
}

func (h HomeController) Create(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "Index",
	})
}

func (h HomeController) Detail(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "Index",
	})
}

func (h HomeController) Update(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "Index",
	})
}

func (h HomeController) Remove(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "Index",
	})
}