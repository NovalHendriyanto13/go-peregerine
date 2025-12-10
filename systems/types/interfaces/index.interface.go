package interfaces

import (
	"github.com/gofiber/fiber/v2"
)

type IndexControllerInterface interface {
	Index(c *fiber.Ctx) error
	Detail(c *fiber.Ctx) error
	Update(c *fiber.Ctx) error
	Create(c *fiber.Ctx) error
	Remove(c *fiber.Ctx) error
}
