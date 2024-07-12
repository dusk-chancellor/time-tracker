package http

import (
	"encoding/json"
	"net/http"

	"github.com/dusk-chancellor/time-tracker/models"
)

// @Summary Start or stop a task
// @Description Начать отсчет времени по задаче для пользователя/Закончить отсчет времени по задаче для пользователя
// @ID start-stop-task
// @Accept json
// @Produce json
// @Param action query string true "Action to perform: start or stop"
// @Success 200 {string} string "Task ID"
// @Failure 400 {string} string "Bad request: Invalid input"
// @Failure 500 {string} string "Internal server error"
// @Router /task [post]
func (h *Handlers) StartStopTaskHandler() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		var task models.Task
			if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
				h.logger.Error(err.Error())
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

		action := r.URL.Query().Get("action")
		switch action {
		case startTask:
			taskID, err := h.appService.StartTask(h.ctx, task)
			if err != nil {
				h.logger.Error(err.Error())
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			w.Write([]byte(taskID))
			w.WriteHeader(http.StatusOK)
		case stopTask:
			taskID, err := h.appService.StopTask(h.ctx, task.Name)
			if err != nil {
				h.logger.Error(err.Error())
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			w.Write([]byte(taskID))
			w.WriteHeader(http.StatusOK)
		default:
			h.logger.Error("Invalid 'action' query")
			http.Error(w, "Invalid 'action' query", http.StatusBadRequest)
			return
		}
	}
}

// @Summary Get user worklist
// @Description Получение трудозатрат по пользователю за период задача-сумма часов и минут с сортировкой от большей затраты к меньшей
// @ID get-user-worklist
// @Produce json
// @Param user_id query string false "User ID"
// @Success 200 {object} TaskResponse "User worklist"
// @Failure 406 {string} string "Not acceptable: Invalid user ID"
// @Failure 500 {string} string "Internal server error"
// @Router /task [get]
func (h *Handlers) GetUserWorklistHandler() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		userID := r.URL.Query().Get("user_id")
		if userID == "" {
			userIDCookie, err := r.Cookie(userIDCookie)
			if err != nil {
				h.logger.Error(err.Error())
				http.Error(w, err.Error(), http.StatusNotAcceptable)
				return
			}
			if userIDCookie == nil {
				h.logger.Error("Invalid user id")
				http.Error(w, "Invalid user id", http.StatusNotAcceptable)
				return
			}
			userID = userIDCookie.Value
		}

		tasks, err := h.appService.GetUserWorklist(h.ctx, userID)
		if err != nil {
			h.logger.Error(err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		resp := TaskResponse{User: userID, Tasks: tasks}
		if err := json.NewEncoder(w).Encode(resp); err != nil {
			h.logger.Error(err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
