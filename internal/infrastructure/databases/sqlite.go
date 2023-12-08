package databases

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
	"github.com/yescorihuela/tasks_management/internal/domain/entities"
	"github.com/yescorihuela/tasks_management/internal/domain/repositories"
)

type SQLIteRepository struct {
	db *sql.DB
}

func NewSQLiteDBConnection() (*sql.DB, error) {

	sqlite, err := sql.Open("sqlite3", "internal/infrastructure/databases/sqlite-files/database.db")

	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return sqlite, nil
}

func NewSQLiteRepository(db *sql.DB) repositories.TaskRepository {
	return &SQLIteRepository{
		db: db,
	}
}

func (repo *SQLIteRepository) Save(task entities.Task) (entities.Task, error) {
	return entities.Task{}, nil
}

func (repo *SQLIteRepository) GetById(id string) (entities.Task, error) {
	return entities.Task{}, nil
}

func (repo *SQLIteRepository) GetByName(name string) (entities.Task, error) {
	return entities.Task{}, nil
}

func (repo *SQLIteRepository) Update(id string, task entities.Task) (entities.Task, error) {
	return entities.Task{}, nil
}

func (repo *SQLIteRepository) Delete(id string) error {
	return nil
}
