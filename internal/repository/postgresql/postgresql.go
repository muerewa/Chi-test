package postgresql

import (
	"Chi-test/internal/domain/models"
	"Chi-test/internal/repository"
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
)

type DB struct {
	db *pgx.Conn
}

func New(conn string) (*DB, error) {
	c, err := pgx.Connect(context.Background(), conn)
	if err != nil {
		return nil, err
	}
	return &DB{c}, nil
}

func (db *DB) FetchById(ctx context.Context, id int64) (*models.Task, error) {
	var task models.Task
	err := db.db.QueryRow(ctx, "SELECT * FROM tasks WHERE id = $1", id).Scan(&task)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, repository.ErrTaskNotFound
		}
		return nil, err
	}
	return &task, nil
}

func (db *DB) CreateTask(ctx context.Context, task *models.Task) (int64, error) {
	var id int64
	err := db.db.QueryRow(ctx, "INSERT INTO tasks (title, description, completed, CreatedAt) VALUES ($1, $2, $3, $4) RETURNING id", task.Title, task.Description, task.Completed, task.CreatedAt).Scan(&id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return 0, repository.ErrTaskNotFound
		}
		return 0, err
	}
	return id, nil
}

func (db *DB) UpdateTask(ctx context.Context, task *models.Task) (*models.Task, error) {
	var updatedTask models.Task
	err := db.db.QueryRow(ctx, "UPDATE tasks SET title = $1, description = $2, completed = $3 WHERE id = $4"+
		"RETURNING id, title, description, completed, created_at", task.Title, task.Description, task.Completed, task.ID).
		Scan(&updatedTask.ID, &updatedTask.Title, &updatedTask.Description, &updatedTask.Completed, &updatedTask.CreatedAt)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, repository.ErrTaskNotFound
		}
		return nil, err
	}
	return &updatedTask, nil
}

func (db *DB) DeleteTask(ctx context.Context, id int64) error {
	_, err := db.db.Exec(ctx, "DELETE FROM tasks WHERE id = $1", id)
	if errors.Is(err, pgx.ErrNoRows) {
		return repository.ErrTaskNotFound
	}
	return err
}
