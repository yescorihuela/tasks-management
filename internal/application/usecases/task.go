package usecases

import (
	"github.com/yescorihuela/tasks_management/internal/domain/entities"
	"github.com/yescorihuela/tasks_management/internal/domain/repositories"
)

type TaskUseCase interface {
	Save(task entities.Task) (entities.Task, error)
	GetById(id int) (entities.Task, error)
	GetByName(name string) (entities.Task, error)
	Update(id int, task entities.Task) (entities.Task, error)
	Delete(id int) error
}

type taskUseCase struct {
	taskRepository repositories.TaskRepository
}

func NewTaskUseCase(taskRepository repositories.TaskRepository) TaskUseCase {
	return &taskUseCase{
		taskRepository: taskRepository,
	}
}

func (s *taskUseCase) Save(task entities.Task) (entities.Task, error) {
	return s.taskRepository.Save(task)
}

func (s *taskUseCase) GetById(id int) (entities.Task, error) {
	return s.taskRepository.GetById(id)
}

func (s *taskUseCase) GetByName(name string) (entities.Task, error) {
	return s.taskRepository.GetByName(name)
}

func (s *taskUseCase) Update(id int, task entities.Task) (entities.Task, error) {
	return s.taskRepository.Update(id, task)
}

func (s *taskUseCase) Delete(id int) error {
	return s.taskRepository.Delete(id)
}
