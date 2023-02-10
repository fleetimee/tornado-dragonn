package middleware

import (
	"github.com/gofiber/fiber/v2"
)

func handle404(c *fiber.Ctx) error {
	return c.Status(404).JSON(fiber.Map{
		"message": "Route not found",
	})
}
