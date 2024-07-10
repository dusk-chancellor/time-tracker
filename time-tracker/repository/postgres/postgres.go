package postgres

import (
	"database/sql"
	"fmt"
	"log/slog"

	"github.com/dusk-chancellor/time-tracker/configs"
	_ "github.com/lib/pq"
)

type Storage struct {
	logger *slog.Logger
	db *sql.DB
}

func NewDB(cfg *configs.Config, logger *slog.Logger) (*Storage, error) {
	db, err := sql.Open("postgres",
		fmt.Sprintf(
			"postgresql://%s:%s@%s:%s/%s?sslmode=disable",
			cfg.DB.User,
			cfg.DB.Password,
			cfg.DB.Host,
			cfg.DB.Port,
			cfg.DB.Name,
		),
	)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return &Storage{logger: logger, db: db}, nil
}
