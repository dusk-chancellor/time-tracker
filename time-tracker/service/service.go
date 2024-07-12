package service

import (
	"context"
	"log/slog"
	"time"

	"github.com/dusk-chancellor/time-tracker/models"
	swaggerAPI "github.com/dusk-chancellor/time-tracker/swagger_api"
)

const (
	filterByID		   = "id"
	filterByPassport   = "passport"
	filterBySurname	   = "surname"
	filterByName	   = "name"
	filterByPatronymic = "patronymic"
	filterByAddress	   = "address"
)

type Service struct {
	logger *slog.Logger
	DBMethods DBMethods
	APIClient *swaggerAPI.APIClient
}

type DBMethods interface {
	AddUser(ctx context.Context, user models.User) (int32, error)
	GetUser(ctx context.Context, pSerie, pNumber int32) (models.User, error)
	UpdateUser(ctx context.Context, user models.User) (int32, error)
	DeleteUser(ctx context.Context, pSerie, pNumber int32) error
	GetAllUsers(ctx context.Context) ([]models.User, error)
	CreateTask(ctx context.Context, task models.Task) (int32, error)
	GetTask(ctx context.Context, taskName string) (models.Task, error)
	UpdateTaskStart(ctx context.Context, startTime time.Time, taskName string) (int32, error)
	UpdateTaskEnd(ctx context.Context, endTime, spentTime time.Time, taskName string) (int32, error)
	GetAllTasksByUserID(ctx context.Context, userID int32) ([]models.Task, error)
	TaskExists(ctx context.Context, taskName string) (bool, error)
}

func NewService(logger *slog.Logger, dbMethods DBMethods, apiClient *swaggerAPI.APIClient) *Service {
	return &Service{
		logger: logger,
		DBMethods: dbMethods,
		APIClient: apiClient,
	}
}
