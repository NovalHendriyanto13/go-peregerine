// Package Interfaces is system package to control type of default interface
package interfaces

import (
	"github.com/gofiber/fiber/v2"
)

// IndexControllerInterface is group of function to handle default action in controller
type IndexControllerInterface interface {
	Index(c *fiber.Ctx) error
	Detail(c *fiber.Ctx) error
	Update(c *fiber.Ctx) error
	Create(c *fiber.Ctx) error
	Remove(c *fiber.Ctx) error
}
