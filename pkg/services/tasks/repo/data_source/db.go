package datasource

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

const DbLocation = "/Users/jovan/go-learning/gprc-Demo/gRPCExample/tasks.db"

type DataSource interface {
	DB() *gorm.DB
}

type DataSourceImpl struct {
	db *gorm.DB
}

type Entity struct {
	Id string `gorm:"primaryKey"`
}

func (dataSource *DataSourceImpl) DB() *gorm.DB {
	if dataSource.db == nil {
		db, err := gorm.Open(sqlite.Open(DbLocation), &gorm.Config{})
		if err != nil {
			panic("failed to connect database")
		}
		dataSource.db = db
	}
	return dataSource.db
}

func NewTasksDataSource() *DataSourceImpl {
	return &DataSourceImpl{}
}
