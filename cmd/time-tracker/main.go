package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/dusk-chancellor/time-tracker/app"
	"github.com/dusk-chancellor/time-tracker/configs"
)

const (
	envDev  = "dev"
	envProd = "prod"
)

// TODO: Add user handler		  [+]
// TODO: Edit user handler		  [+]
// TODO: Delete user handler	  [+]
// TODO: Start & Stop handler	  [+]
// TODO: Get user worklist		  [+]
// TODO: Get all users data		  [+]
// TODO: All database methods	  [+]
// TODO: Cover code with logs	  [+]
// TODO: Config .env file		  [+]
// TODO: Generate swagger for API [-]

func main() {
	cfg := configs.ReadConfig()
	logger := initLogger(cfg.Env)
	ctx := context.Background()

	app := app.NewApp(ctx, logger, cfg)
	app.MigrateDB()

	fmt.Printf("Server started at %s\n", cfg.Server.Host+":"+cfg.Server.Port)
	go func() {
		app.Run()
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	<-stop
	app.Shutdown()
}

func initLogger(env string) *slog.Logger {
	var logger *slog.Logger
	switch env {
	case envDev:
		logger = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case envProd:
		logger = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	}
	return logger
}
