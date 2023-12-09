package repositories

import (
	"github.com/yescorihuela/tasks_management/internal/domain/entities"
)

type TaskRepository interface {
	Save(task entities.Task) (entities.Task, error)
	GetById(id string) (entities.Task, error)
	Update(id string, task entities.Task) (entities.Task, error)
	Delete(id string) error
}
