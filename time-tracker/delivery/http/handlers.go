package http

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/dusk-chancellor/time-tracker/models"
)

const (
	startTask = "start"
	stopTask  = "stop"
)

type Handlers struct {
	appService AppService
	ctx 	   context.Context
	logger 	   *slog.Logger
}

type AppService interface {
	CreateUser(ctx context.Context, passport string) (string, error)
	EditUser(ctx context.Context, newUser models.User) (string, error)
	DeleteUser(ctx context.Context, passport string) error 
}

type Passport struct {
	PassportNumber string `json:"passport_number"`
}

func NewHandlers(appService AppService, ctx context.Context, logger *slog.Logger) *Handlers {
	return &Handlers{appService, ctx, logger}
}

func (h *Handlers) AddUserHandler() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		var passport Passport
		if err := json.NewDecoder(r.Body).Decode(&passport); err != nil {
			h.logger.Error(err.Error())
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		userID, err := h.appService.CreateUser(h.ctx, passport.PassportNumber)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		resp := fmt.Sprintf("User created successfully, your id: %s", userID)
		w.Write([]byte(resp))
		w.WriteHeader(http.StatusOK)
	}
}

func (h *Handlers) EditUserHandler() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {	
		var user models.User
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			h.logger.Error(err.Error())
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		userID, err := h.appService.EditUser(h.ctx, user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		resp := fmt.Sprintf("User updated successfully, your id: %s", userID)
		w.Write([]byte(resp))
		w.WriteHeader(http.StatusOK)
	}
}

func (h *Handlers) DeleteUserHandler() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		var passport Passport
		if err := json.NewDecoder(r.Body).Decode(&passport); err != nil {
			h.logger.Error(err.Error())
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if err := h.appService.DeleteUser(h.ctx, passport.PassportNumber); err != nil {
			h.logger.Error(err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write([]byte("User deleted successfully"))
		w.WriteHeader(http.StatusOK)
	}
}

func (h *Handlers) StartStopTaskHandler() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		action := r.URL.Query().Get("action")
		switch action {
		case startTask:

		case stopTask:

		default:
			http.Error(w, "Invalid query type", http.StatusBadRequest)
		}
	}
}
