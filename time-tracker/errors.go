package timetracker

import "errors"

var (
	ErrUserExists 	= errors.New("user already exists")
	ErrUserNotFound = errors.New("user not found")
	ErrUknownFilter = errors.New("unknown filter")
)
