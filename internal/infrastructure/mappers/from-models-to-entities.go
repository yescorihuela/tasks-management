package mappers

import (
	"github.com/yescorihuela/tasks_management/internal/domain/entities"
	"github.com/yescorihuela/tasks_management/internal/infrastructure/models"
)

func TaskModelToTaskEntity(tm models.Task) entities.Task {
	return entities.Task{
		Id:          tm.Id,
		Title:       tm.Title,
		Description: tm.Description,
		ExpiresAt:   tm.ExpiresAt,
		Status:      tm.Status,
	}
}
