package service

import (
	"context"
	"reflect"
	"strconv"
	"strings"

	"github.com/dusk-chancellor/time-tracker/models"
)

func (s *Service) CreateUser(ctx context.Context, passport string) (string, error) {
	p := strings.Split(passport, " ")
	pSerie, _ := strconv.Atoi(p[0])
	pNumber, _ := strconv.Atoi(p[1])
	people, resp, err := s.APIClient.DefaultAPI.InfoGet(ctx).PassportSerie(int32(pSerie)).PassportNumber(int32(pNumber)).Execute()
	if err != nil {
		s.logger.Error("error: %s, response: %d", err.Error(), resp.StatusCode)
		return "", err
	}

	user := models.User{
		PassportSerie:  int32(pSerie),
		PassportNumber: int32(pNumber),
		Surname:        people.Surname,
		Name:           people.Name,
		Patronymic:     *people.Patronymic,
		Address:        people.Address,
	}

	userID, err := s.DBMethods.AddUser(ctx, user)
	if err != nil {
		s.logger.Error(err.Error())
		return "", err
	}

	s.logger.Info("User %s created, id: %s", user.Name, userID)
	return string(userID), nil
}

func (s *Service) EditUser(ctx context.Context, newUser models.User) (string, error) {
	if newUser.PassportSerie == 0 || newUser.PassportNumber == 0 {
		s.logger.Error("Passport number or passport serie is empty")
		return "", nil
	}
	oldUser, err := s.DBMethods.GetUser(ctx, newUser.PassportSerie, newUser.PassportNumber)
	if err != nil {
		s.logger.Error(err.Error())
		return "", err
	}
	updatedUser := mergeUserInfo(oldUser, newUser)

	userID, err := s.DBMethods.UpdateUser(ctx, updatedUser)
	if err != nil {
		s.logger.Error(err.Error())
		return "", err
	}

	return string(userID), nil
}

func (s *Service) DeleteUser(ctx context.Context, passport string) error {
	p := strings.Split(passport, " ")
	pSerie, _ := strconv.Atoi(p[0])
	pNumber, _ := strconv.Atoi(p[1])

	err := s.DBMethods.DeleteUser(ctx, int32(pSerie), int32(pNumber))
	if err != nil {
		s.logger.Error(err.Error())
		return err
	}
	return nil
}

func mergeUserInfo(oldUser models.User, newUser models.User) models.User {
	newReflect := reflect.ValueOf(&newUser).Elem()
	oldReflect := reflect.ValueOf(oldUser)
	for i := 0; i < oldReflect.NumField(); i++ {
		newField := newReflect.Field(i)
		if !newField.IsValid() || newField.Interface() == reflect.Zero(newField.Type()).Interface() {
            newReflect.Field(i).Set(oldReflect.Field(i))
        }
	}
	return newUser
}
