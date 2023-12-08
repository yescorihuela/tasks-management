package repositories

import (
	"github.com/yescorihuela/tasks_management/internal/domain/entities"
)

type TaskRepository interface {
	Save(task entities.Task) (*entities.Task, error)
	GetById(id int) (entities.Task, error)
	GetByName(name string) (entities.Task, error)
	Update(id int, task entities.Task) (entities.Task, error)
	Delete(id int) error
}
