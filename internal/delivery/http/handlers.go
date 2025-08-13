package http

import (
	"context"
	"log/slog"

	"github.com/dusk-chancellor/time-tracker/internal/models"
)

const (
	startTask 		   = "start"
	stopTask		   = "stop"
	userIDCookie	   = "user_id"
	passportCookie	   = "passport_number"
	filterByID		   = "id"
	filterByPassport   = "passport"
	filterBySurname	   = "surname"
	filterByName	   = "name"
	filterByPatronymic = "patronymic"
	filterByAddress	   = "address"
)

type Handlers struct {
	srv 	Service
	ctx 	context.Context
	logger 	*slog.Logger
}

type Service interface {
	CreateUser(ctx context.Context, passport string) (string, error)
	EditUser(ctx context.Context, newUser models.User) (string, error)
	DeleteUser(ctx context.Context, passport string) error
	StartTask(ctx context.Context, task models.Task) (string, error)
	StopTask(ctx context.Context, taskName string) (string, error)
	GetUserWorklist(ctx context.Context, userID string) ([]models.Task, error)
	GetAllUsersData(ctx context.Context, filter, page string) ([]models.User, error)
}

func NewHandlers(srv Service, ctx context.Context, logger *slog.Logger) *Handlers {
	return &Handlers{srv, ctx, logger}
}
