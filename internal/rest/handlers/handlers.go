package handlers

import (
	"Chi-test/internal/domain/models"
	"context"
	"net/http"
)

type TaskService interface {
	GetTask(ctx context.Context, id int64) (*models.Task, error)
	CreateTask(ctx context.Context, task *models.Task) (int64, error)
	UpdateTask(ctx context.Context, id int64, task *models.Task) (*models.Task, error)
	DeleteTask(ctx context.Context, id int64) error
}

func GetTask(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}
func CreateTask(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}
func DeleteTask(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}
func PutTask(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}
