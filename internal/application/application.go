package application

import "github.com/gin-gonic/gin"

type Application struct {
	taskHandler TaskHandler
	router      *gin.Engine
}

func NewApplication(
	taskHandler TaskHandler,
	router *gin.Engine,
) *Application {
	return &Application{
		taskHandler: taskHandler,
		router:      router,
	}
}
