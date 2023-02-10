package handlers

import (
	"github.com/fleetimee/tornado-dragonn/config"
	"github.com/fleetimee/tornado-dragonn/entities"
	"github.com/fleetimee/tornado-dragonn/helper"
	"github.com/gofiber/fiber/v2"
)

// Generate swagger documentation for this function

// @Summary Get all users
// @Description Get all users
// @Tags users
// @Accept json
// @Produce json
// @Param page query integer false "Page number"
// @Param limit query integer false "Limit number"
// @Router /users [get]
func GetUser(c *fiber.Ctx) error {
	page, limit, err := helper.GetPaginationParams(c)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Invalid page or limit number",
		})
	}

	offset := (page - 1) * limit

	var users []entities.User

	config.Database.Limit(limit).Offset(offset).Find(&users)

	// Get total number of users
	var total int64
	config.Database.Model(&entities.User{}).Count(&total)

	// Return the max number of pages

	return c.JSON(
		fiber.Map{
			"message": "Data Fetch Successfully",
			"status":  c.Context().Response.StatusCode(),
			"data":    users,
			"total":   total,
			"pages":   page,
		},
	)

	// return c.Status(200).JSON(fiber.Map{
	// 	"message": "Data Fetch Successfully",
	// 	"status":  c.Context().Response.StatusCode(),
	// 	"data":    user,
	// })

}

// Generate swagger documentation for this function
// @Summary Add a new user
// @Description Add a new user
// @Tags users
// @Accept json
// @Produce json
// @Param user body entities.User true "User object"
// @Router /users [post]
// @Success 200 {object} entities.User
func AddUser(c *fiber.Ctx) error {
	var user entities.User

	if err := c.BodyParser(&user); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Error parsing data",
		})
	}

	// Hash password before saving to database
	hashedPassword, err := helper.HashPassword(user.Password)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Error hashing password",
		})
	}

	user.Password = hashedPassword

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
