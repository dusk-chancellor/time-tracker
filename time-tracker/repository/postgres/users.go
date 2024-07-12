package postgres

import (
	"context"

	"github.com/dusk-chancellor/time-tracker/models"
	timetracker "github.com/dusk-chancellor/time-tracker/time-tracker"
	"github.com/lib/pq"
)


func (s *Storage) AddUser(ctx context.Context, user models.User) (int32, error) {

	query, err := s.db.Prepare(
		`INSERT INTO users (passport_serie, passport_number, surname, name, patronymic, address)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id`,
	)
	if err != nil {
		s.logger.Error(err.Error())
		return 0, err
	}

	var userID int32
	err = query.QueryRowContext(
		ctx, user.PassportSerie, user.PassportNumber, user.Surname, user.Name, user.Patronymic, user.Address,
		).Scan(&userID)
	if err != nil {
		if pgErr, ok := err.(*pq.Error); ok && pgErr.Code.Name() == "unique_violation" {
			s.logger.Error(err.Error())
			return 0, timetracker.ErrUserExists
		}
		s.logger.Error(err.Error())
		return 0, err
	}
	s.logger.Info("User created")
	return userID, nil
}

func (s *Storage) GetUser(ctx context.Context, pSerie, pNumber int32) (models.User, error) {

	query, err := s.db.Prepare(
		`SELECT id, surname, name, patronymic, address WHERE passport_serie = $1 AND passport_number = $2`,
	)
	if err != nil {
		s.logger.Error(err.Error())
		return models.User{}, err
	}

	user := models.User{
		PassportSerie:  pSerie,
		PassportNumber: pNumber,
	}
	err = query.QueryRowContext(
		ctx, pSerie, pNumber,
	).Scan(
		&user.Id,
		&user.Surname,
		&user.Name,
		&user.Patronymic,
		&user.Address,
	)
	if err != nil {
		s.logger.Error(err.Error())
		return models.User{}, err
	}
	s.logger.Info("User fetched")
	return user, nil
}

func (s *Storage) GetAllUsers(ctx context.Context) ([]models.User, error) {

	query, err := s.db.Prepare(
		`SELECT id, passport_serie, passport_number, surname, name, patronymic, address
		FROM users`,
	)
	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}

	rows, err := query.QueryContext(ctx)
	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}

	users := make([]models.User, 0)
	for rows.Next() {
		user := models.User{}
		err = rows.Scan(
			&user.Id,
			&user.PassportSerie,
			&user.PassportNumber,
			&user.Surname,
			&user.Name,
			&user.Patronymic,
			&user.Address,
		)
		if err != nil {
			s.logger.Error(err.Error())
			return nil, err
		}
		users = append(users, user)
	}
	s.logger.Info("All users fetched")
	return users, nil
}

func (s *Storage) UpdateUser(ctx context.Context, user models.User) (int32, error) {

	query, err := s.db.Prepare(
		`UPDATE users
		SET name = $1, surname = $2, patronymic = $3, address = $4
		WHERE passport_serie = $6 AND passport_number = $7
		RETURNING id`,
	)
	if err != nil {
		s.logger.Error(err.Error())
		return 0, err
	}

	var userID int32
	err = query.QueryRowContext(
		ctx,
		user.Name,
		user.Surname,
		user.Patronymic,
		user.Address,
		user.PassportSerie,
		user.PassportNumber,
	).Scan(&userID)
	if err != nil {
		s.logger.Error(err.Error())
		return 0, err
	}
	s.logger.Info("User updated")
	return userID, nil
}

func (s *Storage) DeleteUser(ctx context.Context, pSerie, pNumber int32) error {

	query, err := s.db.Prepare(
		`DELETE FROM users
		WHERE passport_serie = $1 AND passport_number = $2`,
	)
	if err != nil {
		s.logger.Error(err.Error())
		return err
	}

	_, err = query.ExecContext(
		ctx, pSerie, pNumber,
	)
	if err != nil {
		s.logger.Error(err.Error())
		return err
	}

	s.logger.Info("User %s %s deleted", string(pSerie), string(pNumber))
	return nil
}
