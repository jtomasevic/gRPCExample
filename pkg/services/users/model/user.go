package model

import (
	db "github.com/jtomasevic/gRPCExample/pkg/services/users/repo/data_source"
)

// Taks we have to do
type User struct {
	db.Entity
	Email         string
	Password      string
	FirstName     *string
	LastName      *string
	NumberOfTasks int
}
