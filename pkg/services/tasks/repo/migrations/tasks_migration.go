package migrations

import (
	ds "github.com/jtomasevic/gRPCExample/pkg/services/tasks/repo/data_source"
	model "github.com/jtomasevic/gRPCExample/pkg/services/tasks/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Strptr(value string) *string {
	return &value
}

func TasksDbAutoMigrate() {
	db, err := gorm.Open(sqlite.Open(ds.DbLocation), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(model.Task{})
	// if clearUsers {
	// 	db.Where("ID LIKE ?", "%").Delete(task.Task{})
	// 	db.Where("ID LIKE ?", "%").Delete(user.User{})
	// }
}
