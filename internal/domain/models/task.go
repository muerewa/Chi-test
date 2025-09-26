package models

import "time"

type Task struct {
	ID          int64     `json:"id,omitempty" validate:"number,min=1"`
	Title       string    `json:"title" validate:"required,min=1,max=100"`
	Description string    `json:"description" validate:"required,max=500"`
	Completed   bool      `json:"completed" validate:"required,boolean"`
	CreatedAt   time.Time `json:"created_at,omitempty" validate:"datetime"`
}
