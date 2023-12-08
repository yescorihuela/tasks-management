package entities

import (
	"time"

	"github.com/oklog/ulid/v2"
	"github.com/yescorihuela/tasks_management/internal/infrastructure/mappers/validator"
)

func NewId() string {
	return ulid.Make().String()
}

type Task struct {
	Id          string
	Title       string
	Description string
	ExpiresAt   time.Time
	Status      string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func ValidateTask(task Task, v *validator.Validator) {
	v.Check(len(task.Id) > 0, "id", "invalid or not generated")
	v.Check(len(task.Title) > 3 && len(task.Title) <= 255, "title", "invalid or incompleted title")
	v.Check(len(task.Description) > 3 && len(task.Description) <= 255, "description", "invalid description")
	v.Check(v.In(task.Status, []string{"no completado", "completado", "en proceso"}), "status", "invalid status")
}

func ValidateTaskId(id string, v *validator.Validator) {
	v.Check(len(id) == 26, "id", "invalid ulid id")
}
