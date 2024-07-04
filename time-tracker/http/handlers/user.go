package handlers

import (
	"context"
	"net/http"
)

type UserMethods interface {
	AddUser() ()
	EditUser() ()
	DeleteUser() ()
}

func AddUserHandler(ctx context.Context, userMethods UserMethods) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		
	}
}

func EditUserHandler(ctx context.Context, userMethods UserMethods) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func DeleteUserHandler(ctx context.Context, userMethods UserMethods) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

	}
}
