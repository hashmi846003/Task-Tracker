package models

import "time"

type Task struct {
	ID        int       `json:"id"`
	Description string  `json:"description"`
	Status     string   `json:"status"` // "todo", "in-progress", or "done"
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
}
