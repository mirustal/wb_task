package router

import (
	"task-l0/internal/cache"
	"task-l0/internal/router/handlers"

	"task-l0/pkg/configs"
	"task-l0/platform/database"
	nats_s "task-l0/platform/nats"

	"github.com/gofiber/fiber/v2"
)



func Init(app *fiber.App, storage *database.API, cfg *configs.Config, cache *cache.Cache, sc nats_s.Client) {
	app.Post("/orders", func(c *fiber.Ctx) error {
		return handlers.CreateOrder(c, storage, sc, cache)
	})
	app.Get("/orders/:id", handlers.GetOrder(cache))
	app.Post("/publish", handlers.PublishOrder(sc))
}
