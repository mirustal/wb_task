package app

import (
	"context"
	"log/slog"
	"task-l0/internal/cache"
	"task-l0/internal/router"
	"task-l0/pkg/configs"
	"task-l0/pkg/logging"
	"task-l0/platform/database"
	nats_s "task-l0/platform/nats"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/nats-io/stan.go"
)

type App struct {
	App *fiber.App
	Storage *database.API
	Config *configs.Config
	Log *slog.Logger
	SC stan.Conn
}

const (
	clusterID = "order-nats"
	clientID  = "order-client"
	channel   = "order-channel"
)

func InitApp(cfg *configs.Config, log *logging.Logger) *App {
	db, err := database.NewDatabase(&cfg.PostgresDB)
	if err != nil {
		log.Logger.Error("failed to coonect database", logging.Err(err))
		return &App{}
	}

	err = db.MigrateDB(&cfg.PostgresDB)
	if err != nil {
		log.Logger.Error("failed to migrate database", logging.Err(err))
		return &App{}
	}
	
	cache := cache.NewCache()
	storage := database.NewApi(db.GetDB(), cache)
	storage.LoadAndCacheOrders(context.Background())

	app := fiber.New()
	app.Use(logger.New())
	sc, err := nats_s.NewClient(context.TODO(), cfg)

	router.Init(app, storage, cfg, cache, sc)

	return &App{
		App: app,
		Storage: storage,
		Config: cfg,
		Log: log.Logger,
	}
}

func (app *App) Run() {
	if err := app.App.Listen(app.Config.Http.Port); err != nil {
		app.Log.Error("Server not running", logging.Err(err))
	}
}