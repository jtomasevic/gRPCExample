package repo

import (
	"github.com/google/uuid"
	model "github.com/jtomasevic/gRPCExample/pkg/services/tasks/model"
)

func (repo *tasksRepositoryImpl) Save(entity model.Task) (model.Task, error) {
	if entity.Id == "" {
		entity.Id = uuid.NewString()
		repo.db.Create(entity)
	} else {
		repo.db.Save(entity)
	}
	return entity, nil
}

func (repo *tasksRepositoryImpl) GetByID(id string) (model.Task, error) {
	var entity model.Task
	repo.db.First(&entity, "Id = ?", id)
	return entity, nil
}

func (repo *tasksRepositoryImpl) Delete(id string) error {
	var entity model.Task
	repo.db.Delete(&entity, "Id = ?", id)
	return nil
}

func (repo *tasksRepositoryImpl) All() ([]model.Task, error) {
	var entities []model.Task
	repo.db.Find(&entities)
	return entities, nil
}
