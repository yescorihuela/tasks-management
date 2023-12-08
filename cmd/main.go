package main

import (
	"github.com/gin-gonic/gin"

	_ "github.com/lib/pq"
	"github.com/yescorihuela/tasks_management/internal/application"
	"github.com/yescorihuela/tasks_management/internal/application/usecases"
	"github.com/yescorihuela/tasks_management/internal/infrastructure/databases"
	"github.com/yescorihuela/tasks_management/internal/infrastructure/http/api"
	"github.com/yescorihuela/tasks_management/internal/infrastructure/repositories"
)

func main() {

	db, err := databases.NewPostgresqlDBConnection()
	if err != nil {
		panic(err)
	}

	defer db.Close()

	postgresTaskRepo := repositories.NewPostgresqlRepository(db)
	taskUsecase := usecases.NewTaskUseCase(postgresTaskRepo)
	handler := api.NewTaskHandler(taskUsecase)
	app := application.NewApplication(handler, gin.Default())

	if err := app.Run(); err != nil {
		panic(err)
	}

}
