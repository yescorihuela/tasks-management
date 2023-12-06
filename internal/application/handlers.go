package application

import taskRepository "github.com/yescorihuela/tasks-management/internal/domain/repositories"

type TaskHandler struct {
	taskRepo taskRepository.TaskRepository
}

func NewApplication(
	taskRepo taskRepository.TaskRepository,
) *TaskHandler {
	return &TaskHandler{
		taskRepo: taskRepo,
	}
}

func (app *TaskHandler) Save() {}

func (app *TaskHandler) GetById() {}

func (app *TaskHandler) GetByName() {}

func (app *TaskHandler) Update() {}

func (app *TaskHandler) Delete() {}
