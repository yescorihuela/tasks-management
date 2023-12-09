package application

import (
	"github.com/gin-gonic/gin"
	"github.com/yescorihuela/tasks_management/internal/infrastructure/http/api"
)

type Application struct {
	taskHandler *api.TaskHandler
	router      *gin.Engine
}

func NewApplication(
	taskHandler *api.TaskHandler,
	router *gin.Engine,
) *Application {
	return &Application{
		taskHandler: taskHandler,
		router:      router,
	}
}

func (app *Application) RegisterRoutes() {
	app.router.POST("/task/new", app.taskHandler.Save)
	app.router.GET("/task/by-id/:id", app.taskHandler.GetById)
	app.router.PATCH("task/:id", app.taskHandler.Update)
	app.router.DELETE("task/:id", app.taskHandler.Delete)
}

func (app *Application) Bootstrapping() {
	app.RegisterRoutes()
}

func (app *Application) Run() error {
	app.Bootstrapping()
	err := app.router.Run(":8000")
	return err
}
