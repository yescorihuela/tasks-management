package models

import "time"

type Task struct {
	Id          int       `json:"id"`
	Title       string    `json:"task"`
	Description string    `json:"description"`
	ExpiresAt   time.Time `json:"expires_at"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
