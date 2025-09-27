package service

import (
	"Chi-test/internal/domain/models"
	"context"
	"log/slog"

	"github.com/go-playground/validator/v10"
)

type TaskRepository interface {
	FetchById(ctx context.Context, id int64) (*models.Task, error)
	CreateTask(ctx context.Context, task *models.Task) (int64, error)
	UpdateTask(ctx context.Context, id int64, task *models.Task) (*models.Task, error)
	DeleteTask(ctx context.Context, id int64) error
}

type TaskService struct {
	repo      TaskRepository
	log       *slog.Logger
	validator *validator.Validate
}

func NewTaskService(repo TaskRepository, log *slog.Logger) *TaskService {
	return &TaskService{
		repo:      repo,
		log:       log,
		validator: validator.New(),
	}
}

func (s *TaskService) GetTask(ctx context.Context, id int64) (*models.Task, error) {
	task, err := s.repo.FetchById(ctx, id)
	if err != nil {
		s.log.Error("Failed to fetch task", "id", id, "err", err)
		return nil, err
	}
	return task, nil
}

func (s *TaskService) CreateTask(ctx context.Context, task *models.Task) (int64, error) {
	err := s.validator.Struct(task)
	if err != nil {
		s.log.Info("Failed to validate task", "task", task)
		s.log.Error("validator err", "err", err)
		return 0, err
	}
	res, err := s.repo.CreateTask(ctx, task)
	if err != nil {
		s.log.Error("create err", "err", err)
		return 0, err
	}
	return res, nil
}

func (s *TaskService) UpdateTask(ctx context.Context, id int64, task *models.Task) (*models.Task, error) {
	err := s.validator.Struct(task)
	if err != nil {
		s.log.Error("validator err", "err", err)
		return nil, err
	}
	res, err := s.repo.UpdateTask(ctx, id, task)
	if err != nil {
		s.log.Error("update err", "err", err)
		return nil, err
	}
	return res, nil
}

func (s *TaskService) DeleteTask(ctx context.Context, id int64) error {
	err := s.repo.DeleteTask(ctx, id)
	if err != nil {
		s.log.Error("delete err", "err", err)
	}
	return err
}
