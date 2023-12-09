package utils

import (
	"time"

	"github.com/yescorihuela/tasks_management/internal/domain/entities"
)

type inputTaskPartialUpdate struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
	ExpiresAt   *string `json:"expires_at"`
	Status      *string `json:"status"`
	CreatedAt   *string `json:"created_at"`
	UpdatedAt   *string `json:"updated_at"`
}

func NewInputPartialUpdate() inputTaskPartialUpdate {
	return inputTaskPartialUpdate{}
}

func (i *inputTaskPartialUpdate) BuildTaskToPartialUpdate(task entities.Task) (entities.Task, int) {
	qtyUpdatedFields := 0
	now := time.Now()
	if i.Title != nil && task.Title != *i.Title {
		task.Title = *i.Title
		qtyUpdatedFields++
	}

	if i.Description != nil && task.Description != *i.Description {
		task.Description = *i.Description
		qtyUpdatedFields++
	}

	if i.ExpiresAt != nil {
		expiresAt, err := time.Parse("2006-01-02 15:04:05", *i.ExpiresAt)
		if err != nil {
			return task, 0
		}
		if task.ExpiresAt.UTC() != expiresAt {
			task.ExpiresAt = expiresAt
			qtyUpdatedFields++
		}
	}

	if i.Status != nil && task.Status != *i.Status {
		task.Status = *i.Status
		qtyUpdatedFields++
	}
	task.UpdatedAt = now
	return task, qtyUpdatedFields
}
