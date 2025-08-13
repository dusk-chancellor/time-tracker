package postgres

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/dusk-chancellor/time-tracker/configs"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Storage struct {
	logger *slog.Logger
	pool *pgxpool.Pool
}

func NewDB(cfg *configs.Config, logger *slog.Logger) (*Storage, error) {
	connStr := fmt.Sprintf(
		"postgresql://%s:%s@%s:%s/%s?sslmode=disable&pool_max_conns=10&pool_max_conn_lifetime=1h30m",
		cfg.DB.User,
		cfg.DB.Password,
		cfg.DB.Host,
		cfg.DB.Port,
		cfg.DB.Name,
	)

	dbCfg, err := pgxpool.ParseConfig(connStr)
	if err != nil {
		return nil, err
	}

	pool, err := pgxpool.NewWithConfig(context.Background(), dbCfg)
	if err != nil {
		return nil, err
	}

	return &Storage{
		logger: logger,
		pool: pool,
	}, err
}
