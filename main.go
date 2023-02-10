package main

import (
	"github.com/gofiber/swagger"

	"github.com/fleetimee/tornado-dragonn/config"
	"github.com/fleetimee/tornado-dragonn/handlers"
	"github.com/gofiber/fiber/v2"

	_ "github.com/fleetimee/tornado-dragonn/docs"
	jwtware "github.com/gofiber/jwt/v3"
)

// @title Tornado Dragonn API Documentation
// @description This is the documentation for the Tornado Dragonn API.
// @version 1.0.0
// @schemes http https
// @host localhost:3000
// @BasePath /
func main() {
	app := fiber.New()

	config.Connect()

	// Placeholders
	app.Get("/access", accessible)

	app.Post("/login", handlers.Login)

	app.Get("/swagger/*", swagger.HandlerDefault)

	app.Get("/swagger/*", swagger.New(swagger.Config{ // custom
		URL:         "http://localhost:3000/swagger/doc.json", // The url pointing to API definition
		DeepLinking: false,
		// Expand ("list") or Collapse ("none") tag groups by default
		DocExpansion: "none",
		// Prefill OAuth ClientId on Authorize popup
		OAuth: &swagger.OAuthConfig{
			AppName:  "OAuth Provider",
			ClientId: "21bb4edc-05a7-4afc-86f1-2e151e4ba6e2",
		},
		// Ability to change OAuth2 redirect uri location
		OAuth2RedirectUrl: "http://localhost:8080/swagger/oauth2-redirect.html",
	}))

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "fleetime",
		})
	})

	app.Get("/users", handlers.GetUser)
	app.Post("/users", handlers.AddUser)
	app.Put("/users/:id", handlers.UpdateUser)
	app.Delete("/users/:id", handlers.RemoveUser)

	// Handle unknown routes
	app.Use(func(c *fiber.Ctx) error {
		return c.Status(404).JSON(fiber.Map{
			"message": "Route not found",
			"status":  c.Context().Response.StatusCode(),
		})
	})

	app.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte("secret"),
	}))

	app.Listen(":3000")
}

// Placeholders
func accessible(c *fiber.Ctx) error {
	return c.SendString("Accessible")
}
