package app

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/dusk-chancellor/time-tracker/configs"
	"github.com/dusk-chancellor/time-tracker/time-tracker/repository/postgres"
	"github.com/dusk-chancellor/time-tracker/time-tracker/service"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	swaggerAPI "github.com/dusk-chancellor/time-tracker/swagger_api"
	handlers "github.com/dusk-chancellor/time-tracker/time-tracker/delivery/http"
)

type App struct {
	HttpServer http.Server
	logger 	   *slog.Logger
	cfg 	   *configs.Config
	// ...
}

func NewApp(ctx context.Context, logger *slog.Logger, cfg *configs.Config) *App {
	db, err := postgres.NewDB(cfg, logger)
	if err != nil {
		panic(err)
	}

	apiClient := swaggerAPI.NewAPIClient(swaggerAPI.NewConfiguration())

	appService := service.NewService(logger, db, apiClient)

	appHandlers := handlers.NewHandlers(appService, ctx, logger)

	mux := http.NewServeMux()
	mux.HandleFunc("POST   /user", appHandlers.AddUserHandler())
	mux.HandleFunc("PATCH  /user", appHandlers.EditUserHandler())
	mux.HandleFunc("REMOVE /user", appHandlers.DeleteUserHandler())
	mux.HandleFunc("POST   /task", appHandlers.StartStopTaskHandler())

	app := &App{
		HttpServer: http.Server{
			Addr:    fmt.Sprintf(":%s", cfg.Server.Port),
			Handler: mux,
		},
		logger: logger,
		cfg:    cfg,
	}
	return app
}


func (a *App) Run() error {
	return a.HttpServer.ListenAndServe()
}

func (a *App) Shutdown() {
	if err := a.HttpServer.Shutdown(context.Background()); err != nil {
		panic(err)
	}
}

func (a *App) MigrateDB() {
	m, err := migrate.New(
		"file://"+a.cfg.MigrationsPath,
		fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable",
			a.cfg.DB.User,
			a.cfg.DB.Password,
			a.cfg.DB.Host,
			a.cfg.DB.Port,
			a.cfg.DB.Name,
		),
	)
	if err != nil {
		panic(err)
	}

	if err = m.Up(); err != nil {
		if errors.Is(err, migrate.ErrNoChange) {
			a.logger.Info("no migrations transformed")
			return
		}
		panic(err)
	}

	a.logger.Info("migrations transformed")
}
