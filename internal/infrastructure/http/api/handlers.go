package api

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yescorihuela/tasks_management/internal/application/usecases"
	"github.com/yescorihuela/tasks_management/internal/domain/entities"
	"github.com/yescorihuela/tasks_management/internal/infrastructure/mappers"
	"github.com/yescorihuela/tasks_management/internal/infrastructure/mappers/validator"
	"github.com/yescorihuela/tasks_management/internal/infrastructure/models"
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
		return
	}

	entityTask, err := mappers.FromRequestToEntity(req)

	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"errors": err})
		return
	}

	savedTask, errSave := app.taskUseCase.Save(entityTask)
	taskResponse := mappers.FromEntityToResponse(savedTask)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": errSave})
		return
	}
	ctx.JSON(http.StatusCreated, taskResponse)
}

func (app *TaskHandler) GetById(ctx *gin.Context) {
	id := ctx.Param("id")
	validator := validator.NewValidator()
	if entities.ValidateTaskId(id, validator); !validator.IsValid() {
		ctx.JSON(http.StatusNotFound, gin.H{"errors": validator.Errors})
		return
	}
	task, err := app.taskUseCase.GetById(id)
	taskResponse := mappers.FromEntityToResponse(task)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, taskResponse)
}

func (app *TaskHandler) GetByName(ctx *gin.Context) {}

func (app *TaskHandler) Update(ctx *gin.Context) {}

func (app *TaskHandler) Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	validator := validator.NewValidator()
	if entities.ValidateTaskId(id, validator); !validator.IsValid() {
		ctx.JSON(http.StatusNotFound, gin.H{"errors": validator.Errors})
		return
	}
	err := app.taskUseCase.Delete(id)

	if errors.Is(models.ErrorRecordNotFound, err) {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	} else if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}
