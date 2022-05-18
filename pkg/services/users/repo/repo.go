package repo

import "github.com/google/uuid"
import model "github.com/jtomasevic/gRPCExample/pkg/services/users/model"


func (repo *usersRepositoryImpl) Save(entity model.User) (model.User, error) {
	if entity.Id == "" {
		entity.Id = uuid.NewString()
		repo.db.Create(entity)
	} else {
		repo.db.Save(entity)
	}
	return entity, nil
}

func (repo *usersRepositoryImpl) GetByID(id string) (model.User, error) {
	var entity model.User
	repo.db.First(&entity, "Id = ?", id)
	return entity, nil
}

func (repo *usersRepositoryImpl) Delete(id string) error {
	var entity model.User
	repo.db.Delete(&entity, "Id = ?", id)
	return nil
}

func (repo *usersRepositoryImpl) All() ([]model.User, error) {
	var entities []model.User
	repo.db.Find(&entities)
	return entities, nil
}

func (repo *usersRepositoryImpl) GetByUserName(userName string) (model.User, error) {
	var entity model.User
	repo.db.First(&entity, "UserName = ?", userName)
	return entity, nil
}
