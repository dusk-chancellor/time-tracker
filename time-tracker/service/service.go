package service

import (
	"context"
	"log/slog"

	"github.com/dusk-chancellor/time-tracker/models"
	swaggerAPI "github.com/dusk-chancellor/time-tracker/swagger_api"
)

type Service struct {
	logger *slog.Logger
	DBMethods UserDBMethods
	APIClient *swaggerAPI.APIClient
}

type UserDBMethods interface {
	AddUser(ctx context.Context, user models.User) (int32, error)
	GetUser(ctx context.Context, pSerie, pNumber int32) (models.User, error)
	UpdateUser(ctx context.Context, user models.User) (int32, error)
	DeleteUser(ctx context.Context, pSerie, pNumber int32) error
	GetAllUsers(ctx context.Context) ([]models.User, error)
}

func NewService(logger *slog.Logger, dbMethods UserDBMethods, apiClient *swaggerAPI.APIClient) *Service {
	return &Service{
		logger: logger,
		DBMethods: dbMethods,
		APIClient: apiClient,
	}
}