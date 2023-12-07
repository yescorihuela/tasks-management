package databases

import (
	"github.com/jmoiron/sqlx"
	"github.com/yescorihuela/tasks_management/internal/domain/entities"
	"github.com/yescorihuela/tasks_management/internal/domain/repositories"
)

type PostgresqlRepository struct {
	db *sqlx.DB
}

func NewPostgresqlDBConnection() (*sqlx.DB, error) {
	return nil, nil
}

func NewPostgresqlRepository(db *sqlx.DB) repositories.TaskRepository {
	return &PostgresqlRepository{
		db: db,
	}
}

func (repo *PostgresqlRepository) Save(task entities.Task) (entities.Task, error) {
	return entities.Task{}, nil
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
