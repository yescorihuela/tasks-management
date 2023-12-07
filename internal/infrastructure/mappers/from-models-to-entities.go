package infrastructure

import (
	taskEntity "github.com/yescorihuela/tasks_management/internal/domain/entities"
	taskModel "github.com/yescorihuela/tasks_management/internal/infrastructure/models"
)

func TaskModelToTaskEntity(tm taskModel.Task) taskEntity.Task {
	return taskEntity.Task{
		Id:          tm.Id,
		Title:       tm.Title,
		Description: tm.Description,
		ExpiresAt:   tm.ExpiresAt,
		Status:      tm.Status,
	}
}
