package postgresql

import (
	"database/sql"

	taskEntity "github.com/yescorihuela/tasks-management/internal/domain/entities"
)

type PostgreqlRepository struct {
	db *sql.DB
}

func NewPostgreqlRepository(db *sql.DB) *PostgreqlRepository {
	return &PostgreqlRepository{
		db: db,
	}
}

func (repo *PostgreqlRepository) Save(task taskEntity.Task) (taskEntity.Task, error) {
	return taskEntity.Task{}, nil
}

func (repo *PostgreqlRepository) GetById(id int) (taskEntity.Task, error) {
	return taskEntity.Task{}, nil
}

func (repo *PostgreqlRepository) GetByName(id int) (taskEntity.Task, error) {
	return taskEntity.Task{}, nil
}

func (repo *PostgreqlRepository) Update(id int, task taskEntity.Task) (taskEntity.Task, error) {
	return taskEntity.Task{}, nil
}

func (repo *PostgreqlRepository) Delete(id int) error {
	return nil
}
