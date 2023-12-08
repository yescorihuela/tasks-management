package repositories

import (
	"database/sql"
	"time"

	"github.com/yescorihuela/tasks_management/internal/domain/entities"
	"github.com/yescorihuela/tasks_management/internal/domain/repositories"
	"github.com/yescorihuela/tasks_management/internal/infrastructure/models"
)

type PostgresqlRepository struct {
	db *sql.DB
}

func NewPostgresqlRepository(db *sql.DB) repositories.TaskRepository {
	return &PostgresqlRepository{
		db: db,
	}
}

func (repo *PostgresqlRepository) Save(task entities.Task) (*entities.Task, error) {
	queryPostgres := `
		INSERT INTO TASKS(title, description, status, created_at, updated_at)
		VALUES($1, $2, $3, $4, $5) RETURNING id, title, description, status
	`
	currentDate := time.Now()
	var modelTask models.Task
	err := repo.db.QueryRow(queryPostgres, task.Title, task.Description, task.Status, currentDate, currentDate).
		Scan(
			&modelTask.Id,
			&modelTask.Title,
			&modelTask.Description,
			&modelTask.Status,
		)
	if err != nil {
		return nil, err
	}

	newTask := entities.Task{
		Id:          modelTask.Id,
		Title:       modelTask.Title,
		Description: modelTask.Description,
		Status:      modelTask.Status,
	}

	return &newTask, nil
}

func (repo *PostgresqlRepository) GetById(id int) (entities.Task, error) {
	return entities.Task{}, nil
}

func (repo *PostgresqlRepository) GetByName(name string) (entities.Task, error) {
	return entities.Task{}, nil
}

func (repo *PostgresqlRepository) Update(id int, task entities.Task) (entities.Task, error) {
	return entities.Task{}, nil
}

func (repo *PostgresqlRepository) Delete(id int) error {
	return nil
}
