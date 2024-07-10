package models

type User struct {
	Id        	   int32
	PassportSerie  int32  `json:"passport_serie"`
	PassportNumber int32  `json:"passport_number"`
	Surname 	   string `json:"surname"`
	Name 		   string `json:"name"`
	Patronymic 	   string `json:"patronymic"`
	Address 	   string `json:"address"`
}