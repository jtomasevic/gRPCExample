package repo

import (
	"gorm.io/gorm"
	db "github.com/jtomasevic/gRPCExample/pkg/services/tasks/repo/data_source"
	model "github.com/jtomasevic/gRPCExample/pkg/services/tasks/model"
)

type TasksRepository interface {
	Save(entity model.Task) (model.Task, error)
	GetByID(id string) (model.Task, error)
	Delete(id string) error
	All() ([]model.Task, error)
}

type tasksRepositoryImpl struct {
	db   gorm.DB
}

func NewTaskRepository(dataSource db.DataSource) *tasksRepositoryImpl {
	return &tasksRepositoryImpl{
		db: *dataSource.DB(),
	}
}

