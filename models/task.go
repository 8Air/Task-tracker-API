package models

import "time"

const (
	StatusNew        = "new"
	StatusInProgress = "in_progress"
	StatusDone       = "done"
)

type Task struct {
	ID          int       `json:"id" swaggerignore:"true"`
	Title       string    `json:"title"`
	Description string    `json:"description,omitempty"`
	Status      string    `json:"status" enums:"new,in_progress,done" example:"new"`
	CreatedAt   time.Time `json:"created_at" example:"2025-04-22T14:00:00Z" swaggerignore:"true"`
	UpdatedAt   time.Time `json:"updated_at" example:"2025-04-22T14:00:00Z" swaggerignore:"true"`
}
