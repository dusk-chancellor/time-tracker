package app

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/dusk-chancellor/time-tracker/configs"
	"github.com/dusk-chancellor/time-tracker/internal/repository"
	"github.com/dusk-chancellor/time-tracker/internal/repository/postgres"
	"github.com/dusk-chancellor/time-tracker/internal/service"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	handlers "github.com/dusk-chancellor/time-tracker/internal/delivery/http"
	swaggerAPI "github.com/dusk-chancellor/time-tracker/swagger_api"
)

type App struct {
	HttpServer http.Server
	logger 	   *slog.Logger
	cfg 	   *configs.Config
	// ...
}

func NewApp(ctx context.Context, logger *slog.Logger, cfg *configs.Config) *App {
	pool, err := postgres.ConnectDB(cfg, logger)
	if err != nil {
		panic(err)
	}

	migrateDB(logger, cfg)

	apiCfg := &swaggerAPI.Configuration{
		DefaultHeader: make(map[string]string),
		Debug:         false,
		Servers: []swaggerAPI.ServerConfiguration{
			{
				URL: cfg.OuterAPI,
				Description: "Outer API",
			},
		},
		OperationServers: map[string]swaggerAPI.ServerConfigurations{
		},
	}

	repo := repository.NewRepo(logger, pool)
	apiClient := swaggerAPI.NewAPIClient(apiCfg)
	appService := service.NewService(logger, repo, repo, apiClient)
	appHandlers := handlers.NewHandlers(appService, ctx, logger)

	mux := http.NewServeMux()

	mux.HandleFunc("GET    /user", appHandlers.GetAllUsersDataHandler)
	mux.HandleFunc("POST   /user", appHandlers.AddUserHandler)
	mux.HandleFunc("PATCH  /user", appHandlers.EditUserHandler)
	mux.HandleFunc("DELETE /user", appHandlers.DeleteUserHandler)

	mux.HandleFunc("POST   /task", appHandlers.StartStopTaskHandler)
	mux.HandleFunc("GET    /task", appHandlers.GetUserWorklistHandler)

	app := &App{
		HttpServer: http.Server{
			Addr:    fmt.Sprintf(":%s", cfg.ServerPort),
			Handler: mux,
		},
		logger: logger,
		cfg:    cfg,
	}
	return app
}


func (a *App) Run() error {
	if err := a.HttpServer.ListenAndServe(); err != nil {
		if errors.Is(err, http.ErrServerClosed) {
			a.logger.Info("server closed")
			return err
		} 
		a.logger.Error(err.Error())
		return err
	}

	return nil
}

func (a *App) Shutdown(ctx context.Context) {
	a.HttpServer.Shutdown(ctx)
}

func migrateDB(l *slog.Logger, cfg *configs.Config) {
	m, err := migrate.New(
		cfg.MigrationsPath,
		cfg.DBUrl,
	)
	if err != nil {
		panic(err)
	}

	if err = m.Up(); err != nil {
		if errors.Is(err, migrate.ErrNoChange) {
			l.Info("no migrations transformed")
			return
		}
		panic(err)
	}
	l.Info("migrations transformed")
}
