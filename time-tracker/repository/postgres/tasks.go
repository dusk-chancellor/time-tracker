package postgres

import (
	"context"
	"time"

	"github.com/dusk-chancellor/time-tracker/models"
)


func (s *Storage) CreateTask(ctx context.Context, task models.Task) (int32, error) {
	query, err := s.db.Prepare(
		`INSERT INTO tasks (user_id, name, description, start_time, created_at)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id`,
	)
	if err != nil {
		s.logger.Error(err.Error())
		return 0, err
	}

	var taskID int32
	err = query.QueryRowContext(
		ctx, task.UserId, task.Name, task.Description, task.StartTime, task.CreatedAt,
	).Scan(&taskID)
	if err != nil {
		s.logger.Error(err.Error())
		return 0, err
	}
	return taskID, nil
}

func (s *Storage) UpdateTaskStart(ctx context.Context, startTime time.Time, taskName string) (int32, error) {
	query, err := s.db.Prepare(
		`UPDATE tasks
		SET start_time = $1
		WHERE name = $2
		RETURNING id`,
	)
	if err != nil {
		s.logger.Error(err.Error())
		return 0, err
	}

	var taskID int32
	err = query.QueryRowContext(ctx, startTime, taskName).Scan(&taskID)
	if err != nil {
		s.logger.Error(err.Error())
		return 0, err
	}
	return taskID, nil
}

func (s *Storage) GetTask(ctx context.Context, taskName string) (models.Task, error) {
	query, err := s.db.Prepare(
		`SELECT *
		FROM tasks
		WHERE name = $1`,
	)
	if err != nil {
		s.logger.Error(err.Error())
		return models.Task{}, err
	}

	var task models.Task
	err = query.QueryRowContext(
		ctx, taskName,
	).Scan(&task)
	if err != nil {
		s.logger.Error(err.Error())
		return models.Task{}, err
	}
	return task, nil
}

func (s *Storage) UpdateTaskEnd(ctx context.Context, endTime, spentTime time.Time, taskName string) (int32, error) {
	query, err := s.db.Prepare(
		`UPDATE tasks
		SET end_time = $1, spent_time = $2
		WHERE name = $3
		RETURNING id`,
	)
	if err != nil {
		s.logger.Error(err.Error())
		return 0, err
	}

	var taskID int32
	err = query.QueryRowContext(ctx, endTime, spentTime, taskName).Scan(&taskID)
	if err != nil {
		s.logger.Error(err.Error())
		return 0, err
	}
	return taskID, nil
}

func (s *Storage) GetAllTasksByUserID(ctx context.Context, userID int32) ([]models.Task, error) {
	query, err := s.db.Prepare(
		`SELECT *
		FROM tasks
		WHERE user_id = $1`,
	)
	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}

	var tasks []models.Task
	err = query.QueryRowContext(
		ctx, userID,
	).Scan(&tasks)
	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}
	return tasks, nil
}

func (s *Storage) TaskExists(ctx context.Context, taskName string) (bool, error) {
    var count int
    query := `SELECT COUNT(*) FROM tasks WHERE name = $1`
    err := s.db.QueryRowContext(ctx, query, taskName).Scan(&count)
    if err != nil {
		s.logger.Error(err.Error())
        return false, err
    }
    return count > 0, nil
}
