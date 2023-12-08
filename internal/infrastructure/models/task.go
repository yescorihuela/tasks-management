package models

import (
	"errors"
	"time"

	"github.com/yescorihuela/tasks_management/internal/domain/entities"
)

var (
	ErrorRecordNotFound = errors.New("record not found")
)

type Task struct {
	Id          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	ExpiresAt   time.Time `json:"expires_at"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func AsModel(entity entities.Task) Task {
	now := time.Now()
	return Task{
		Id:          entity.Id,
		Title:       entity.Title,
		Description: entity.Description,
		Status:      entity.Status,
		ExpiresAt:   entity.ExpiresAt,
		CreatedAt:   now,
		UpdatedAt:   now,
	}
}
