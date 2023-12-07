package application

import (
	"github.com/gin-gonic/gin"
	"github.com/yescorihuela/tasks_management/internal/infrastructure/http/api"
)

type Application struct {
	taskHandler api.TaskHandler
	router      *gin.Engine
}

func NewApplication(
	taskHandler api.TaskHandler,
	router *gin.Engine,
) *Application {
	return &Application{
		taskHandler: taskHandler,
		router:      router,
	}
}
