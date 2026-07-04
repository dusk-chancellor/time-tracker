package repository

import (
	"log/slog"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Repo struct {
	l *slog.Logger
	p  *pgxpool.Pool
}

func NewRepo(logger *slog.Logger, pool *pgxpool.Pool) *Repo {
	return &Repo{
		l: logger,
		p: pool,
	}
}