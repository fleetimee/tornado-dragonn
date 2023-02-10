package middleware

import (
	"github.com/gofiber/fiber/v2"

	"github.com/golang-jwt/jwt/v4"
)

func restricted(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["name"].(string)
	return c.SendString("Welcome " + name)
}
