package main

import (

	"task-l0/internal/app"
	"task-l0/pkg/configs"
	"task-l0/pkg/logging"
)

func main() {

	// os.Setenv("CONFIG_PATH", "/Users/mirustal/Documents/project/go/wildberries/L0/config.yml")


	cfg := configs.GetConfig()
	log := logging.SetupLogger(cfg.ModeLog)
	// log.Info("Starting service", slog.String("env", cfg.ModeLog))
	app := app.InitApp(cfg, log)
	app.Run()

}