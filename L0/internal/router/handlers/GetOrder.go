package handlers

import (
	"task-l0/internal/cache"

	"github.com/gofiber/fiber/v2"
)


func GetOrder(cache *cache.Cache) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		if id == "" {
			return c.Status(fiber.StatusBadRequest).SendString("ID is required")
		}

		data, found := cache.Get(id)
		if !found {
			return c.Status(fiber.StatusNotFound).SendString("Order not found")
		}

		return c.JSON(data)
	}
}