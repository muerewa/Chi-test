package service

import (
	"Chi-test/internal/domain/models"
	"context"

	"github.com/go-playground/validator/v10"
)

type TaskRepository interface {
	FetchById(ctx context.Context, id int64) (*models.Task, error)
	CreateTask(ctx context.Context, task *models.Task) (int64, error)
	UpdateTask(ctx context.Context, task *models.Task) (*models.Task, error)
	DeleteTask(ctx context.Context, id int64) error
}

type TaskService struct {
	repo      TaskRepository
	validator *validator.Validate
}

func NewTaskService(repo TaskRepository) *TaskService {
	return &TaskService{
		repo:      repo,
		validator: validator.New(),
	}
}
