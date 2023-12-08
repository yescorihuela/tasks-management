package mappers

import (
	"time"

	"github.com/yescorihuela/tasks_management/internal/domain/entities"
)

type TaskRequest struct {
	Id          string `json:"id,omitempty"`
	Title       string `json:"title"`
	Description string `json:"description"`
	ExpiresAt   string `json:"expires_at,omitempty"`
	Status      string `json:"status"`
	CreatedAt   string `json:"created_at,omitempty"`
	UpdatedAt   string `json:"updated_at,omitempty"`
}

type TaskResponse struct {
	Id          string    `json:"id,omitempty"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	ExpiresAt   time.Time `json:"expires_at,omitempty"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
}

func FromEntityToResponse(task entities.Task) TaskResponse {
	return TaskResponse{
		Id:          task.Id,
		Title:       task.Title,
		Description: task.Description,
		ExpiresAt:   task.ExpiresAt.UTC(),
		Status:      task.Status,
		CreatedAt:   task.CreatedAt.UTC(),
		UpdatedAt:   task.UpdatedAt.UTC(),
	}
}
