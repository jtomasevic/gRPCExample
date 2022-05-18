package repo

import (
	db "github.com/jtomasevic/gRPCExample/pkg/services/tasks/repo/data_source"
)

// Taks we have to do
type Task struct {
	// Unique integer identifier of the todo task
	db.Entity
	// Title of the task
	Title string
	// Detail description of the todo task
	Description string

	UserId string
}
