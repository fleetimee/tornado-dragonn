package helper

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// GetPaginationParams returns the pagination parameters (page and limit) from the request query
func GetPaginationParams(c *fiber.Ctx) (int, int, error) {
	page := c.Query("page")
	limit := c.Query("limit")

	if page == "" {
		page = "1"
	}
	if limit == "" {
		limit = "10"
	}

	p, err := strconv.Atoi(page)
	if err != nil {
		return 0, 0, err
	}

	l, err := strconv.Atoi(limit)
	if err != nil {
		return 0, 0, err
	}

	return p, l, nil
}
