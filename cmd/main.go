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
	defaultCfgPath = "./configs/local.yaml"
)

func main() {
	cfgPath, ok := os.LookupEnv("CONFIG_PATH")
	if !ok {
		cfgPath = defaultCfgPath
	}

	cfg, err := configs.LoadConfig(cfgPath)
	if err != nil {
		panic(err)
	}

	logger := initLogger(cfg.Env)
	ctx := context.Background()

	app := app.NewApp(ctx, logger, cfg)

	fmt.Printf("Server started at %s\n", cfg.Server.Host + ":" + cfg.Server.Port)
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
	case envDev:
		logger = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case envProd:
		logger = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	}
	return logger
}
