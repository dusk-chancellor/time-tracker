package service

import (
	"context"
	"sort"
	"strconv"
	"time"

	"github.com/dusk-chancellor/time-tracker/internal/models"
)


func (s *Service) StartTask(ctx context.Context, task *models.Task) (string, error) {
	now := time.Now()
	exists, err := s.TaskRepo.TaskExists(ctx, task.Name)
	if err != nil {
		s.logger.Error(err.Error())
		return "", err
	}
	if exists {
		err := s.TaskRepo.UpdateTaskStart(ctx, now, task.Name)
		if err != nil {
			s.logger.Error(err.Error())
			return "", err
		}
		s.logger.Info("Task updated")
		return strconv.Itoa(int(task.Id)), nil
	}

	taskID, err := s.TaskRepo.CreateTask(ctx, task)
	if err != nil {
		s.logger.Error(err.Error())
		return "", err
	}
	s.logger.Info("Task created")
	return strconv.Itoa(int(taskID)), nil
}

func (s *Service) StopTask(ctx context.Context, taskName string) (string, error) {
	task, err := s.TaskRepo.GetTask(ctx, taskName)
	if err != nil {
		s.logger.Error(err.Error())
		return "", err
	}
	
	now := time.Now()

	err = s.TaskRepo.UpdateTaskEnd(ctx, now, task.SpentTime, taskName)
	if err != nil {
		s.logger.Error(err.Error())
		return "", err
	}

	s.logger.Info("Task updated")
	return string(task.Id), nil
}

func (s *Service) GetUserWorklist(ctx context.Context, userID string) ([]*models.Task, error) {
	userIDInt, _ := strconv.Atoi(userID)
	tasks, err := s.TaskRepo.GetAllTasksByUserID(ctx, int32(userIDInt))
	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}
	if len(tasks) == 0 {
		s.logger.Info("User has no tasks")
		return nil, nil
	}
	
	sort.Slice(tasks, func(i, j int) bool {
		return tasks[i].SpentTime < (tasks[j].SpentTime)
	})

	s.logger.Info("User worklist fetched")
	return tasks, nil
}
