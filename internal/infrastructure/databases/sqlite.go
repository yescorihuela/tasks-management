package databases

import (
	"github.com/jmoiron/sqlx"
	"github.com/yescorihuela/tasks_management/internal/domain/entities"
	"github.com/yescorihuela/tasks_management/internal/domain/repositories"
)

type SQLIteRepository struct {
	db *sqlx.DB
}

func NewSQLiteDBConnection() (*sqlx.DB, error) {
	return nil, nil
}

func NewSQLiteRepository(db *sqlx.DB) repositories.TaskRepository {
	return &SQLIteRepository{
		db: db,
	}
}

func (repo *SQLIteRepository) Save(task entities.Task) (entities.Task, error) {
	return entities.Task{}, nil
}

func (repo *SQLIteRepository) GetById(id int) (entities.Task, error) {
	return entities.Task{}, nil
}

func (repo *SQLIteRepository) GetByName(name string) (entities.Task, error) {
	return entities.Task{}, nil
}

func (repo *SQLIteRepository) Update(id int, task entities.Task) (entities.Task, error) {
	return entities.Task{}, nil
}

func (repo *SQLIteRepository) Delete(id int) error {
	return nil
}
