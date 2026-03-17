// Package api main package to accommodate HomeController
package api

import (
	// "log"
	"github.com/gofiber/fiber/v2"
	"peregerine/systems/types/responses"
	"peregerine/systems/ai"
	"peregerine/app/requests"
)

// HomeController is group of home routes action
type HomeController struct {
	responses.BaseResponse
}

// Index as an action from home routes to get list of data
func (h HomeController) Index(c *fiber.Ctx) error {
	aiChat, _ := ai.Generate("hello")

	resp := h.SuccessResponse(true, fiber.Map{
		"message": aiChat,
	})
	return c.JSON(resp)
}

// Create as an action from home routes to create data
func (h HomeController) Create(c *fiber.Ctx) error {
	var req requests.HomeCreateRequest
	var aiChat string

	// Parse JSON body
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid request body",
		})
	}

	if req.Message == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "message is required",
		})
	}

	c.Set("Content-Type", "text/plain; charset=utf-8")
	c.Set("Transfer-Encoding", "chunked")

	ch := make(chan string)

	go func() {
		_ = ai.StreamGenerate(req.Message, ch)
		close(ch)
	}()

	for token := range ch {
		aiChat += token
	}

	resp := h.SuccessResponse(true, fiber.Map{
		"message": aiChat,
	})
	return c.JSON(resp)
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