package handlers

import (
	"context"
	"net/http"
)

type TaskMethods interface {
	StartTask() ()
	StopTask() ()
}


func StartTaskHandler(ctx context.Context, taskMethods TaskMethods) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func StopTaskHandler(ctx context.Context, taskMethods TaskMethods) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

	}
}
