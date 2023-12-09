package repositories

import (
	"database/sql"
	"errors"
	"regexp"
	"strings"

	"github.com/yescorihuela/tasks_management/internal/domain/entities"
	"github.com/yescorihuela/tasks_management/internal/domain/repositories"
	"github.com/yescorihuela/tasks_management/internal/infrastructure/models"
)

var spaceEater = regexp.MustCompile(`\s+`)

func compact(q string) string {
	return spaceEater.
		ReplaceAllString(strings.TrimSpace(q), " ")
}

type PostgresqlRepository struct {
	db *sql.DB
}

func NewPostgresqlRepository(db *sql.DB) repositories.TaskRepository {
	return &PostgresqlRepository{
		db: db,
	}
}

func (repo *PostgresqlRepository) Save(task entities.Task) (entities.Task, error) {
	queryPostgres := compact(`
		INSERT INTO tasks
			(id, title, description, status, expires_at, created_at, updated_at)
		VALUES
			($1, $2, $3, $4, $5, $6, $7)
		RETURNING *
	`)

	var modelTask models.Task
	err := repo.db.QueryRow(queryPostgres,
		task.Id,
		task.Title,
		task.Description,
		task.Status,
		task.ExpiresAt.UTC(),
		task.CreatedAt.UTC(),
		task.UpdatedAt.UTC(),
	).
		Scan(
			&modelTask.Id,
			&modelTask.Title,
			&modelTask.Description,
			&modelTask.Status,
			&modelTask.ExpiresAt,
			&modelTask.CreatedAt,
			&modelTask.UpdatedAt,
		)
	if err != nil {
		return entities.Task{}, err
	}

	newTask := entities.Task{
		Id:          modelTask.Id,
		Title:       modelTask.Title,
		Description: modelTask.Description,
		Status:      modelTask.Status,
		ExpiresAt:   modelTask.ExpiresAt,
		CreatedAt:   modelTask.CreatedAt,
		UpdatedAt:   modelTask.UpdatedAt,
	}

	return newTask, nil
}

func (repo *PostgresqlRepository) GetById(id string) (entities.Task, error) {
	queryPostgres := compact(`
		SELECT
			id, title, description, status, expires_at, created_at, updated_at
		FROM tasks
			WHERE id = $1
	`)

	var modelTask models.Task
	err := repo.db.QueryRow(queryPostgres, id).
		Scan(
			&modelTask.Id,
			&modelTask.Title,
			&modelTask.Description,
			&modelTask.Status,
			&modelTask.ExpiresAt,
			&modelTask.CreatedAt,
			&modelTask.UpdatedAt,
		)
	if err != nil {
		return entities.Task{}, err
	}

	tasks := entities.Task{
		Id:          modelTask.Id,
		Title:       modelTask.Title,
		Description: modelTask.Description,
		Status:      modelTask.Status,
		ExpiresAt:   modelTask.ExpiresAt,
		CreatedAt:   modelTask.CreatedAt,
		UpdatedAt:   modelTask.UpdatedAt,
	}
	return tasks, nil
}

func (repo *PostgresqlRepository) GetByName(name string) (entities.Task, error) {
	return entities.Task{}, nil
}

func (repo *PostgresqlRepository) Update(id string, task entities.Task) (entities.Task, error) {
	queryPostgres := compact(`
		UPDATE tasks
		SET
		title = $1,
		description = $2,
		status = $3,
		expires_at = $4,
		created_at = $5,
		updated_at = $6
		WHERE id = $7
		RETURNING *
	`)

	var modelTask models.Task
	err := repo.db.QueryRow(queryPostgres,
		task.Title,
		task.Description,
		task.Status,
		task.ExpiresAt.UTC(),
		task.CreatedAt.UTC(),
		task.UpdatedAt.UTC(),
		task.Id,
	).
		Scan(
			&modelTask.Id,
			&modelTask.Title,
			&modelTask.Description,
			&modelTask.Status,
			&modelTask.ExpiresAt,
			&modelTask.CreatedAt,
			&modelTask.UpdatedAt,
		)

	if err != nil && errors.Is(err, sql.ErrNoRows) {
		return entities.Task{}, models.ErrorRecordNotFound
	} else if err != nil {
		return entities.Task{}, err
	}

	if err != nil {
		return entities.Task{}, err
	}
	updatedTask := entities.Task{
		Id:          modelTask.Id,
		Title:       modelTask.Title,
		Description: modelTask.Description,
		Status:      modelTask.Status,
		ExpiresAt:   modelTask.ExpiresAt,
		CreatedAt:   modelTask.CreatedAt,
		UpdatedAt:   modelTask.UpdatedAt,
	}
	return updatedTask, nil
}

func (repo *PostgresqlRepository) Delete(id string) error {
	queryPostgres := compact(`
		DELETE
		FROM tasks
		WHERE id = $1
		RETURNING id`)
	var returnedId string
	err := repo.db.QueryRow(queryPostgres, id).Scan(&returnedId)
	if err != nil && errors.Is(err, sql.ErrNoRows) {
		return models.ErrorRecordNotFound
	} else if err != nil {
		return err
	}
	return nil
}
