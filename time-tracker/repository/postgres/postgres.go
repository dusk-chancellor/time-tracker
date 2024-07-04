package postgres

import (
	"context"
	"database/sql"
	"fmt"

	config "github.com/dusk-chancellor/time-tracker/configs"
)

type Storage struct {
	db *sql.DB
}

func ConnectToDB(ctx context.Context, cfg *config.Config) (*Storage, error) {
	db, err := sql.Open("postgres", fmt.Sprintf(
		"user=%s password=%s dbname=%s host=%s port=%s sslmode=disable",
		cfg.DB.User,
		cfg.DB.Password,
		cfg.DB.Database,
		cfg.DB.Host,
		cfg.DB.Port,
	))
	if err != nil {
		return nil, err
	}

	if err = db.PingContext(ctx); err != nil {
		return nil, err
	}

	return &Storage{db: db}, nil
}
