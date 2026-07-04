package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/dusk-chancellor/time-tracker/internal/app"
	"github.com/dusk-chancellor/time-tracker/configs"
)

const (
	envDev  = "dev"
	envProd = "prod"
)

func main() {
	cfg, err := configs.LoadConfig()
	if err != nil {
		panic(err)
	}

	logger := initLogger(cfg.Env)
	ctx := context.Background()

	app := app.NewApp(ctx, logger, cfg)

	fmt.Printf("Server started at %s\n", cfg.ServerHost + ":" + cfg.ServerPort)
	go func() {
		if err := app.Run(); err != nil {
			panic(err)
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	<-stop
	app.Shutdown(ctx)
}

func initLogger(env string) *slog.Logger {
	var logger *slog.Logger
	switch env {
	case envProd:
		logger = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	default:
		logger = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	}
	return logger
}
