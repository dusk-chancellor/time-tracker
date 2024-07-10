package service

import (
	"reflect"
	"testing"

	"github.com/dusk-chancellor/time-tracker/models"
)

func TestMergeUserInfo(t *testing.T) {
	cases := []struct {
		name string
		oldUser  models.User
		newUser  models.User
		expected models.User
	} {
		{
			name: "merge user info",
			oldUser: models.User{
				PassportSerie:  123,
				PassportNumber: 123,
				Surname:        "surname",
				Name:           "name",
				Patronymic:     "",
				Address:        "address",
			},
			newUser: models.User{
				Patronymic:     "patronymic",
				Address:        "new address",
			},
			expected: models.User{
				PassportSerie:  123,
				PassportNumber: 123,
				Surname:        "surname",
				Name:           "name",
				Patronymic:     "patronymic",
				Address:        "new address",
			},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			updated := mergeUserInfo(c.oldUser, c.newUser)
			if !reflect.DeepEqual(updated, c.expected) {
				t.Errorf("expected %v, got %v", c.expected, c.oldUser)
			}
		})
	}
}
