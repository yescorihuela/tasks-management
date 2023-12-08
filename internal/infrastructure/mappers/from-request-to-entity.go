package mappers

import (
	"time"

	"github.com/yescorihuela/tasks_management/internal/domain/entities"
	"github.com/yescorihuela/tasks_management/internal/infrastructure/mappers/validator"
)

func FromRequestToEntity(tm TaskRequest) (entities.Task, map[string]string) {
	now := time.Now().UTC()
	newId := entities.NewId()

	validator := validator.NewValidator()
	expiresAt, err := time.Parse("2006-01-02 15:04:05", tm.ExpiresAt)
	if err != nil {
		validator.Check(err == nil, "expires_at", err.Error())
	}
	task := entities.Task{
		Id:          newId,
		Title:       tm.Title,
		Description: tm.Description,
		Status:      tm.Status,
		ExpiresAt:   expiresAt,
		CreatedAt:   now,
		UpdatedAt:   now,
	}
	if entities.ValidateTask(task, validator); !validator.IsValid() {
		return entities.Task{}, validator.Errors
	}
	return task, nil
}
