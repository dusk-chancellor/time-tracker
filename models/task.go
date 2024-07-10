package models

type Task struct {
	Id        	int32
	UserId    	int32
	Name      	string
	Description string
	StartTime 	string
	EndTime   	string
	CreatedAt 	string
	UpdatedAt 	string
}