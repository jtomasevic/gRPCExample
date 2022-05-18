package datasource

import (
	model "github.com/jtomasevic/gRPCExample/pkg/services/users/model"
	ds "github.com/jtomasevic/gRPCExample/pkg/services/users/repo/data_source"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Strptr(value string) *string {
	return &value
}

func UsersDbAutoMigrate() {
	db, err := gorm.Open(sqlite.Open(ds.DbLocation), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(model.User{})
}
