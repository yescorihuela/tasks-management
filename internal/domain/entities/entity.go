package domain

import "time"

type Task struct {
	Id          int
	Title       string
	Description string
	ExpiresAt   time.Time
	Status      string
}

func NewTask(
	id int,
	title, description string,
	expiresAt time.Time,
	status string,
) Task {
	return Task{
		Id:          id,
		Title:       title,
		Description: description,
		ExpiresAt:   expiresAt,
		Status:      status,
	}
}
