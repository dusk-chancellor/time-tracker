package repository

import (
	"context"
	"time"

	"github.com/dusk-chancellor/time-tracker/internal/models"
)


func (r *Repo) CreateTask(ctx context.Context, task *models.Task) (int32, error) {
	query := `INSERT INTO tasks (user_id, name, description, start_time, created_at)
			  VALUES ($1, $2, $3, $4, $5)
			  RETURNING id`

	row := r.p.QueryRow(
		ctx, query,
		task.UserId, task.Name,
		task.Description,
		task.StartTime,
		task.CreatedAt,
	)
	err := row.Scan(&task.Id)
	if err != nil {
		return 0, err
	}

	return task.Id, nil
}

func (r *Repo) UpdateTaskStart(ctx context.Context, startTime time.Time, taskName string) error {
	query := `UPDATE tasks
			  SET start_time = $1
			  WHERE name = $2`

	res, err := r.p.Exec(ctx, query, startTime, taskName)
	if err != nil {
		return err
	}

	if res.RowsAffected() == 0 {
		return ErrNotFound
	}
	return nil
}

func (r *Repo) GetTask(ctx context.Context, taskName string) (*models.Task, error) {
	query := `SELECT id, user_id, name,
					  description, start_time, end_time,
					  created_at, updated_at, spent_time
			  FROM tasks
			  WHERE name = $1`

	var task models.Task
	row := r.p.QueryRow(ctx, query, taskName)
	err := row.Scan(
		&task.Id,
		&task.UserId,
		&task.Name,
		&task.Description,
		&task.StartTime,
		&task.EndTime,
		&task.CreatedAt,
		&task.UpdatedAt,
		&task.SpentTime,
	)
	if err != nil {
		return nil, err
	}

	return &task, nil
}

func (r *Repo) UpdateTaskEnd(ctx context.Context, endTime time.Time, spentTime time.Duration, taskName string) error {
	query := `UPDATE tasks
			  SET end_time = $1, spent_time = $2
			  WHERE name = $3`

	res, err := r.p.Exec(
		ctx, query,
		endTime,
		spentTime,
		taskName,
	)
	if err != nil {
		return err
	}

	if res.RowsAffected() == 0 {
		return ErrNotFound
	}
	return nil
}

func (r *Repo) GetAllTasksByUserID(ctx context.Context, userID int32) ([]*models.Task, error) {
	query := `SELECT id, user_id, name,
					  description, start_time, end_time,
					  created_at, updated_at, spent_time
			  FROM tasks
			  WHERE user_id = $1`

	rows, err := r.p.Query(ctx, query, userID)
	if err != nil {
		return nil, err
	}
	// iterating over rows and appending each task into tasks slice
	var tasks = make([]*models.Task, 0)
	for rows.Next() {
		task := models.Task{}
		err = rows.Scan(
			&task.Id,
			&task.UserId,
			&task.Name,
			&task.Description,
			&task.StartTime,
			&task.EndTime,
			&task.CreatedAt,
			&task.UpdatedAt,
			&task.SpentTime,
		)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, &task)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return tasks, nil
}

func (r *Repo) TaskExists(ctx context.Context, taskName string) (bool, error) {
    query := `SELECT EXISTS
			  (SELECT 1 FROM tasks WHERE name = $1)`

	var exists = false
    row := r.p.QueryRow(ctx, query, taskName)
	err := row.Scan(&exists)
    if err != nil {
        return false, err
    }

    return exists, nil
}
