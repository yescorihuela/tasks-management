package mappers

import (
	"time"

	"github.com/yescorihuela/tasks_management/internal/domain/entities"
)

func FromRequestToEntity(tm TaskRequest) (entities.Task, error) {
	now := time.Now().UTC()
	newId := entities.NewId()
	expiresAt, err := time.Parse("2006-01-02 15:04:05", tm.ExpiresAt)
	if err != nil {
		return entities.Task{}, err
	}
	return entities.Task{
		Id:          newId,
		Title:       tm.Title,
		Description: tm.Description,
		Status:      tm.Status,
		ExpiresAt:   expiresAt,
		CreatedAt:   now,
		UpdatedAt:   now,
	}, nil
}
