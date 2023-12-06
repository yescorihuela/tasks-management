package repositories

import (
	taskEntity "github.com/yescorihuela/tasks-management/internal/domain/entities"
)

type TaskRepository interface {
	Save(task taskEntity.Task) (taskEntity.Task, error)
	GetById(id int) (taskEntity.Task, error)
	GetByName(name string) (taskEntity.Task, error)
	Update(id int, task taskEntity.Task) (taskEntity.Task, error)
	Delete(id int) error
}
