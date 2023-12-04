package usecases

import (
	taskEntity "github.com/yescorihuela/tasks-management/internal/domain/entities"
	taskRepository "github.com/yescorihuela/tasks-management/internal/domain/repositories"
)

type TaskUseCase interface {
	Save(task taskEntity.Task) error
	GetById(id int) (taskEntity.Task, error)
	GetByName(name string) (taskEntity.Task, error)
	Update(id int, task taskEntity.Task) (taskEntity.Task, error)
	Delete(id int) error
}

type taskUseCase struct {
	taskRepository taskRepository.TaskRepository
}

func NewTaskUseCase(taskRepository taskRepository.TaskRepository) TaskUseCase {
	return &taskUseCase{
		taskRepository: taskRepository,
	}
}

func (s *taskUseCase) Save(task taskEntity.Task) error {
	return s.taskRepository.Save(task)
}

func (s *taskUseCase) GetById(id int) (taskEntity.Task, error) {
	return s.taskRepository.GetById(id)
}

func (s *taskUseCase) GetByName(name string) (taskEntity.Task, error) {
	return s.taskRepository.GetByName(name)
}

func (s *taskUseCase) Update(id int, task taskEntity.Task) (taskEntity.Task, error) {
	return s.taskRepository.Update(id, task)
}

func (s *taskUseCase) Delete(id int) error {
	return s.taskRepository.Delete(id)
}
