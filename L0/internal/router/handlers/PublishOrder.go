package handlers

import (
	"encoding/json"
	"fmt"
	"task-l0/internal/models"
	nats_s "task-l0/platform/nats"

	"github.com/gofiber/fiber/v2"
)

func PublishOrder(sc nats_s.Client) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		order := new(models.OrderDTO)
		if err := c.BodyParser(&order); err != nil {
			return c.Status(400).JSON(fiber.Map{
				"error": "invalid body",
			})
		}

		fmt.Println("i am publish", order)
		orderJSON, _ := json.Marshal(order)
		if err := sc.Publish("order-channel", orderJSON); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to publish order",
			})
		}
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Order published successfully",
		})
	}
}
