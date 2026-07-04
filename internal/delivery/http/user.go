package http

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/dusk-chancellor/time-tracker/internal/models"
)

type PassportRequest struct {
	PassportNumber string `json:"passport_number"`
}

func (h *Handlers) GetAllUsersDataHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	ctx := r.Context()

	filter := r.URL.Query().Get("filter")
	page := r.URL.Query().Get("page")

	data, err := h.srv.GetAllUsersData(ctx, filter, page)
	if err != nil {
		h.logger.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err = json.NewEncoder(w).Encode(data); err != nil {
		h.logger.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}


func (h *Handlers) AddUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	ctx := r.Context()

	var passport PassportRequest
	if err := json.NewDecoder(r.Body).Decode(&passport); err != nil {
		h.logger.Error(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	userID, err := h.srv.CreateUser(ctx, passport.PassportNumber)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{Name: "user_id", Value: userID})
	resp := fmt.Sprintf("ID: %s", userID)
	w.Write([]byte(resp))
}

// fix:
func (h *Handlers) EditUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	ctx := r.Context()

	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		h.logger.Error(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err := h.srv.EditUser(ctx, &user)
	if err != nil {
		h.logger.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp := fmt.Sprintf("User %s edited successfully", user.Name)
	w.Write([]byte(resp))
}

func (h *Handlers) DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	ctx := r.Context()

	var passport PassportRequest
	if err := json.NewDecoder(r.Body).Decode(&passport); err != nil {
		h.logger.Error(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.srv.DeleteUser(ctx, passport.PassportNumber); err != nil {
		h.logger.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
