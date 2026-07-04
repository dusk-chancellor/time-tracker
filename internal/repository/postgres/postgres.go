package postgres

import (
	"context"
	"log/slog"
	"time"

	"github.com/dusk-chancellor/time-tracker/configs"
	"github.com/jackc/pgx/v5/pgxpool"
)

func ConnectDB(cfg *configs.Config, l *slog.Logger) (*pgxpool.Pool, error) {
	// build pool cfg
	poolCfg, err := pgxpool.ParseConfig(cfg.DBUrl)
	if err != nil {
		l.Error(err.Error())
		return nil, err
	}

	poolCfg.MinConns = 2
	poolCfg.MaxConns = 10
	poolCfg.MaxConnLifetime = time.Hour
	poolCfg.HealthCheckPeriod = time.Minute
	// set up pool
	var ctx = context.Background()
	pool, err := pgxpool.NewWithConfig(ctx, poolCfg)
	if err != nil {
		l.Error(err.Error())
		return nil, err
	}

	if err := pool.Ping(ctx); err != nil {
		pool.Close()
		l.Error(err.Error())
		return nil, err
	}

	l.Info("succesful db pool connection")
	return pool, err
}
