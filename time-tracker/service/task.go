package service

import (
	"context"
	"sort"
	"strconv"
	"time"

	"github.com/dusk-chancellor/time-tracker/models"
)


func (s *Service) StartTask(ctx context.Context, task models.Task) (string, error) {
	now := time.Now()
	exists, err := s.DBMethods.TaskExists(ctx, task.Name)
	if err != nil {
		s.logger.Error(err.Error())
		return "", err
	}
	if exists {
		taskID, err := s.DBMethods.UpdateTaskStart(ctx, now, task.Name)
		if err != nil {
			s.logger.Error(err.Error())
			return "", err
		}
		s.logger.Info("Task updated")
		return string(taskID), nil
	}

	taskID, err := s.DBMethods.CreateTask(ctx, task)
	if err != nil {
		s.logger.Error(err.Error())
		return "", err
	}
	s.logger.Info("Task created")
	return string(taskID), nil
}

func (s *Service) StopTask(ctx context.Context, taskName string) (string, error) {
	task, err := s.DBMethods.GetTask(ctx, taskName)
	if err != nil {
		s.logger.Error(err.Error())
		return "", err
	}
	
	now := time.Now()
	spentTime := task.SpentTime.Add(now.Sub(task.StartTime))

	taskID, err := s.DBMethods.UpdateTaskEnd(ctx, now, spentTime, taskName)
	if err != nil {
		s.logger.Error(err.Error())
		return "", err
	}

	s.logger.Info("Task updated")
	return string(taskID), nil
}

func (s *Service) GetUserWorklist(ctx context.Context, userID string) ([]models.Task, error) {
	userIDInt, _ := strconv.Atoi(userID)
	tasks, err := s.DBMethods.GetAllTasksByUserID(ctx, int32(userIDInt))
	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}
	if len(tasks) == 0 {
		s.logger.Info("User has no tasks")
		return nil, nil
	}
	
	sort.Slice(tasks, func(i, j int) bool {
		return tasks[i].SpentTime.After(tasks[j].SpentTime)
	})

	s.logger.Info("User worklist fetched")
	return tasks, nil
}
