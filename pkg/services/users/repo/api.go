package repo

import (
	"gorm.io/gorm"
	model "github.com/jtomasevic/gRPCExample/pkg/services/users/model"
	db "github.com/jtomasevic/gRPCExample/pkg/services/users/repo/data_source"
)

type UsersRepository interface {
	Save(entity model.User) (model.User, error)
	GetByID(id string) (model.User, error)
	Delete(id string) error
	All() ([]model.User, error)
	GetByUserName(userName string) (model.User, error)
}

type usersRepositoryImpl struct {
	db   gorm.DB
}

func NewUserRepository(dataSource db.DataSource) *usersRepositoryImpl {
	return &usersRepositoryImpl{
		db: *dataSource.DB(),
	}
}

