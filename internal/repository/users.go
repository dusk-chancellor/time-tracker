package repository

import (
	"context"
	"errors"
	"strconv"

	"github.com/dusk-chancellor/time-tracker/internal/models"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgconn"
)


func (r *Repo) AddUser(ctx context.Context, user *models.User) (int32, error) {
	query := `INSERT INTO users (passport_serie, passport_number, surname, name, patronymic, address)
			  VALUES ($1, $2, $3, $4, $5, $6)
			  RETURNING id`

	var userID int32
	row := r.p.QueryRow(
		ctx, query,
		user.PassportSerie,
		user.PassportNumber,
		user.Surname,
		user.Name,
		user.Patronymic,
		user.Address,
	)
	err := row.Scan(&userID)
	if err != nil {
		var e *pgconn.PgError
		if errors.As(err, &e) && e.Code == pgerrcode.UniqueViolation {
			return 0, ErrUserExists
		}
		return 0, err
	}
	
	r.l.Debug("User created")
	return userID, nil
}

func (r *Repo) GetUser(ctx context.Context, pSerie, pNumber int32) (*models.User, error) {
	query := `SELECT id, surname, name, patronymic, address
			  WHERE passport_serie = $1 AND passport_number = $2
			  FROM users`

	user := &models.User{
		PassportSerie:  pSerie,
		PassportNumber: pNumber,
	}
	row := r.p.QueryRow(
		ctx, query,
		pSerie,
		pNumber,
	)
	err := row.Scan(
		&user.Id,
		&user.Surname,
		&user.Name,
		&user.Patronymic,
		&user.Address,
	)
	if err != nil {
		return &models.User{}, err
	}

	r.l.Debug("User fetched")
	return user, nil
}

func (r *Repo) GetAllUsers(ctx context.Context) ([]*models.User, error) {
	query := `SELECT id, passport_serie, passport_number, surname, name, patronymic, address
			  FROM users`

	rows, err := r.p.Query(ctx, query)
	if err != nil {
		return nil, err
	}

	users := make([]*models.User, 0)
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
			return nil, err
		}
		users = append(users, &user)
	}

	r.l.Debug("All users fetched")
	return users, nil
}

func (r *Repo) UpdateUser(ctx context.Context, user *models.User) (int32, error) {
	query := `UPDATE users
			  SET name = $1, surname = $2, patronymic = $3, address = $4
			  WHERE passport_serie = $5 AND passport_number = $6
			  RETURNING id`

	var userID int32
	row := r.p.QueryRow(
		ctx, query,
		user.Name,
		user.Surname,
		user.Patronymic,
		user.Address,
		user.PassportSerie,
		user.PassportNumber,
	)
	err := row.Scan(&userID)
	if err != nil {
		return 0, err
	}

	r.l.Debug("User updated")
	return userID, nil
}

func (r *Repo) DeleteUser(ctx context.Context, pSerie, pNumber int32) error {
	query := `DELETE FROM users
			  WHERE passport_serie = $1 AND passport_number = $2`

	_, err := r.p.Exec(
		ctx, query,
		pSerie, pNumber,
	)
	if err != nil {
		return err
	}

	r.l.Debug("User %s %s deleted",
		strconv.Itoa(int(pSerie)),
		strconv.Itoa(int(pNumber)),
	)
	return nil
}
