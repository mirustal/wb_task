package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"task-l0/internal/cache"
	"task-l0/internal/models"
	"task-l0/platform/database"
	nats_s "task-l0/platform/nats"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/nats-io/stan.go"
)


func CreateOrder(c *fiber.Ctx, storage *database.API, sc nats_s.Client, cache *cache.Cache) error {
	fmt.Println("CreateOrder start \n", sc)
	var orderData models.OrderDTO
	sub, err := sc.Subscribe("order-channel", func(msg *stan.Msg) {
		if err := json.Unmarshal(msg.Data, &orderData); err != nil {
			fmt.Printf("Unmarshal order %v", orderData)
			return
		}
		fmt.Println("i am subcribe", orderData)

	}, stan.DeliverAllAvailable(), stan.DurableName("my-durable-subscription"))
	time.Sleep(time.Second * 1)
	sub.Close()
	fmt.Println("Connect end")
	if err != nil {
		fmt.Println("Subscribe error %v", err)
		return nil
	}
	if orderData.OrderUid == "" {
		return c.Status(400).JSON(fiber.Map{
			"message": "queue clear",
		})
}

	err = storage.CreateOrder(context.Background(), orderData)
	if err != nil {
			return c.Status(400).JSON(fiber.Map{
				"message": "dublciate order",
			})
	}


	cache.Set(orderData.OrderUid, orderData)
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Order created successfully",
		"data": orderData,
	})
}