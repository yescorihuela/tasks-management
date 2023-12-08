package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yescorihuela/tasks_management/internal/application/usecases"
	"github.com/yescorihuela/tasks_management/internal/domain/entities"
	"github.com/yescorihuela/tasks_management/internal/infrastructure/mappers"
)

type TaskHandler struct {
	taskUseCase usecases.TaskUseCase
}

func NewTaskHandler(
	taskUseCase usecases.TaskUseCase,
) *TaskHandler {
	return &TaskHandler{
		taskUseCase: taskUseCase,
	}
}

func (app *TaskHandler) Save(ctx *gin.Context) {
	var req mappers.TaskRequest
	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	var entityTask entities.Task
	entityTask.Title = req.Title
	entityTask.Description = req.Description
	entityTask.Status = req.Status

	t, err := app.taskUseCase.Save(entityTask)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, t)
}

func (app *TaskHandler) GetById(ctx *gin.Context) {}

func (app *TaskHandler) GetByName(ctx *gin.Context) {}

func (app *TaskHandler) Update(ctx *gin.Context) {}

func (app *TaskHandler) Delete(ctx *gin.Context) {}
