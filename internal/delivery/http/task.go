package http

import (
	"encoding/json"
	"net/http"

	"github.com/dusk-chancellor/time-tracker/internal/models"
)

type TaskResponse struct {
	User  string `json:"user"`
	Tasks []models.Task `json:"tasks"`
}


func (h *Handlers) StartStopTaskHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var task models.Task
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		h.logger.Error(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	action := r.URL.Query().Get("action")
	switch action {
	case startTask:
		taskID, err := h.srv.StartTask(h.ctx, task)
		if err != nil {
			h.logger.Error(err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write([]byte(taskID))
		w.WriteHeader(http.StatusOK)
	case stopTask:
		taskID, err := h.srv.StopTask(h.ctx, task.Name)
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

func (h *Handlers) GetUserWorklistHandler(w http.ResponseWriter, r *http.Request) {
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

	tasks, err := h.srv.GetUserWorklist(h.ctx, userID)
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
