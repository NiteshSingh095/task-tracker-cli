package model

/// Task represents a task with ID, description, status, created and updated timestamps
type Task struct {
	ID int `json:"id"`
	Description string `json:"description"`
	Status string `json:"status"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

/// Define allowed status values as constants
const (
	StatusTodo = "TODO"
	StatusInProgress = "IN_PROGRESS"
	StatusDone = "DONE"
)