package http

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/dusk-chancellor/time-tracker/models"
)

// @Summary Get all users' data
// @Description Получение данных пользователей: Фильтрация по всем полям. Пагинация.
// @ID get-all-users-data
// @Produce json
// @Param filter query string false "Optional filter criteria"
// @Param page query string false "Optional page number for pagination"
// @Success 200 {string} string "Data received successfully"
// @Failure 500 {string} string "Internal server error"
// @Router /user [get]
func (h *Handlers) GetAllUsersDataHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		filter := r.URL.Query().Get("filter")
		page := r.URL.Query().Get("page")

		data, err := h.appService.GetAllUsersData(h.ctx, filter, page)
		if err != nil {
			h.logger.Error(err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write([]byte("Data received successfully"))
		if err = json.NewEncoder(w).Encode(data); err != nil {
			h.logger.Error(err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

// @Summary Add a new user
// @Description Добавление нового пользователя в формате
// @ID add-user
// @Accept json
// @Produce json
// @Param passport body PassportRequest true "Passport information"
// @Success 200 {string} string "User created successfully"
// @Failure 400 {string} string "Bad request: Invalid input"
// @Failure 500 {string} string "Internal server error"
// @Router /user [post]
func (h *Handlers) AddUserHandler() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var passport PassportRequest
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

		http.SetCookie(w, &http.Cookie{Name: "user_id", Value: userID})
		resp := fmt.Sprintf("User created successfully, your id: %s", userID)
		w.Write([]byte(resp))
		w.WriteHeader(http.StatusOK)
	}
}

// @Summary Edit user details
// @Description Изменение данных пользователя
// @ID edit-user
// @Accept json
// @Produce json
// @Param user body models.User true "User object to update"
// @Success 200 {string} string "User edited successfully"
// @Failure 400 {string} string "Bad request: Invalid input"
// @Failure 500 {string} string "Internal server error"
// @Router /user [patch]
func (h *Handlers) EditUserHandler() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var user models.User
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			h.logger.Error(err.Error())
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		_, err := h.appService.EditUser(h.ctx, user)
		if err != nil {
			h.logger.Error(err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		resp := fmt.Sprintf("User %s edited successfully", user.Name)
		w.Write([]byte(resp))
		w.WriteHeader(http.StatusOK)
	}
}

// @Summary Delete a user
// @Description Удаление пользователя
// @ID delete-user
// @Accept json
// @Produce json
// @Param passport body PassportRequest true "Passport information"
// @Success 200 {string} string "User deleted successfully"
// @Failure 400 {string} string "Bad request: Invalid input"
// @Failure 500 {string} string "Internal server error"
// @Router /user [delete]
func (h *Handlers) DeleteUserHandler() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var passport PassportRequest
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
