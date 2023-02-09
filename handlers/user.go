package handlers

import (
	"github.com/fleetimee/tornado-dragonn/config"
	"github.com/fleetimee/tornado-dragonn/entities"
	"github.com/gofiber/fiber/v2"
)

func GetUser(c *fiber.Ctx) error {
	var user []entities.User

	config.Database.Find(&user)
	return c.Status(200).JSON(user)
}

func AddUser(c *fiber.Ctx) error {
	var user entities.User

	if err := c.BodyParser(&user); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Error parsing data",
		})
	}

	config.Database.Create(&user)
	return c.Status(200).JSON(user)
}

func UpdateUser(c *fiber.Ctx) error {
	user := new(entities.User)
	id := c.Params("id")

	if err := c.BodyParser(user); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Error parsing data",
		})
	}

	config.Database.Model(&user).Where("id = ?", id).Updates(user)
	return c.Status(200).JSON(user)
}

func RemoveUser(c *fiber.Ctx) error {
	id := c.Params("id")

	config.Database.Delete(&entities.User{}, id)
	return c.Status(200).JSON(fiber.Map{
		"message": "User deleted",
	})
}
