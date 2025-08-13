package models

import "time"

type Task struct {
	Id        	int32
	UserId    	int32     `json:"user_id"`
	Name      	string    `json:"name"`
	Description string    `json:"description,omitempty"`
	StartTime 	time.Time
	EndTime   	time.Time
	CreatedAt 	time.Time
	UpdatedAt 	time.Time
	SpentTime 	time.Time
}