package entities

import (
	"time"

	"github.com/oklog/ulid/v2"
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
