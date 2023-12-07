package databases

import (
	"github.com/jmoiron/sqlx"
	taskEntity "github.com/yescorihuela/tasks_management/internal/domain/entities"
)

type PostgresqlRepository struct {
	db *sqlx.DB
}

func NewDBConnection() (*sqlx.DB, error) {
	return nil, nil
}

func NewPostgresqlRepository(db *sqlx.DB) *PostgresqlRepository {
	return &PostgresqlRepository{
		db: db,
	}
}

func (repo *PostgresqlRepository) Save(task taskEntity.Task) (taskEntity.Task, error) {
	return taskEntity.Task{}, nil
}

func (repo *PostgresqlRepository) GetById(id int) (taskEntity.Task, error) {
	return taskEntity.Task{}, nil
}

func (repo *PostgresqlRepository) GetByName(id int) (taskEntity.Task, error) {
	return taskEntity.Task{}, nil
}

func (repo *PostgresqlRepository) Update(id int, task taskEntity.Task) (taskEntity.Task, error) {
	return taskEntity.Task{}, nil
}

func (repo *PostgresqlRepository) Delete(id int) error {
	return nil
}
