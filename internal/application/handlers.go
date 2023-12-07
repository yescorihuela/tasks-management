package application

import (
	"github.com/gin-gonic/gin"
	domain "github.com/yescorihuela/tasks_management/internal/domain/entities"
	taskUseCase "github.com/yescorihuela/tasks_management/internal/infrastructure/usecases"
)

type TaskHandler struct {
	taskUseCase taskUseCase.TaskUseCase
}

func NewTaskHandler(
	taskUseCase taskUseCase.TaskUseCase,
) *TaskHandler {
	return &TaskHandler{
		taskUseCase: taskUseCase,
	}
}

func (app *TaskHandler) Save(ctx *gin.Context) {
	app.taskUseCase.Save(domain.Task{})
}

func (app *TaskHandler) GetById(ctx *gin.Context) {}

func (app *TaskHandler) GetByName(ctx *gin.Context) {}

func (app *TaskHandler) Update(ctx *gin.Context) {}

func (app *TaskHandler) Delete(ctx *gin.Context) {}
