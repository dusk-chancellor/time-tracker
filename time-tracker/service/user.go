package service

import (
	"context"
	"reflect"
	"sort"
	"strconv"
	"strings"

	"github.com/dusk-chancellor/time-tracker/models"
	timetracker "github.com/dusk-chancellor/time-tracker/time-tracker"
)

func (s *Service) GetAllUsersData(ctx context.Context, filter, page string) ([]models.User, error) {
	users, err := s.DBMethods.GetAllUsers(ctx)
	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}
	if users == nil {
		return nil, nil
	}
	filteredUsers := filterFunc(users, filter)
	if filteredUsers == nil {
		return nil, timetracker.ErrUknownFilter
	}

	var pagedUsers [][]models.User
	chunkSize := 10
	for i := 0; i < len(filteredUsers); i += chunkSize {
		end := i + chunkSize
		if end > len(filteredUsers) {
			end = len(filteredUsers)
		}
		pagedUsers = append(pagedUsers, filteredUsers[i:end])
	}

	if page == "" {
		return pagedUsers[0], nil
	}
	pageInt, err := strconv.Atoi(page)
	if err != nil {
		return nil, err
	}
	return pagedUsers[pageInt], nil
}

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

	s.logger.Info("User %s %s updated", newUser.Name, newUser.Surname)
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

	s.logger.Info("User %s %s deleted", p[0], p[1])
	return nil
}

func filterFunc(users []models.User, filter string) []models.User {
	switch filter {
	case filterByID:
		sort.Slice(users, func(i, j int) bool {
			return users[i].Id > users[j].Id
		})
	case filterByPassport:
		sort.Slice(users, func(i, j int) bool {
			return users[i].PassportSerie > users[j].PassportSerie
		})
	case filterBySurname:
		sort.Slice(users, func(i, j int) bool {
			return users[i].Surname > users[j].Surname
		})
	case filterByName:
		sort.Slice(users, func(i, j int) bool {
			return users[i].Name > users[j].Name
		})
	case filterByPatronymic:
		sort.Slice(users, func(i, j int) bool {
			return users[i].Patronymic > users[j].Patronymic
		})
	case filterByAddress:
		sort.Slice(users, func(i, j int) bool {
			return users[i].Address > users[j].Address
		})
	case "":
		// nichego ne delat
	default:
		return nil
	}
	return users
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
