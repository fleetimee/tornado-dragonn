package handlers

import (
	"github.com/fleetimee/tornado-dragonn/config"
	"github.com/fleetimee/tornado-dragonn/entities"
	"github.com/fleetimee/tornado-dragonn/helper"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func Login(c *fiber.Ctx) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	// Throw error if username or password is missing or empty
	if username == "" || password == "" {
		return c.Status(400).JSON(fiber.Map{
			"message": "Missing username or password",
		})
	}

	// Check if user exists in database
	var user entities.User
	config.Database.Where("username = ?", username).First(&user)

	// Throw error if user does not exist
	if user.ID == 0 {
		return c.Status(400).JSON(fiber.Map{
			"message": "User does not exist",
		})
	}

	// Check if username and password match
	if !helper.CheckPasswordHash(password, user.Password) {
		return c.Status(400).JSON(fiber.Map{
			"message": "Invalid credentials",
		})
	}

	// Create JWT Claims
	claims := jwt.MapClaims{
		"username": user.Username,
	}

	// Create JWT Token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("secret"))

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Internal Server Error",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Login successful",
		"token":   t,
	})

}
