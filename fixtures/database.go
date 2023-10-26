package fixtures

import (
	"pethost/config"
	"pethost/frameworks/database/gorm_adapter"
	"pethost/frameworks/database/gorm_adapter/models"
)

var TestDatabase = gorm_adapter.New(config.TestDatabase)

var tables = []string{
	models.Pet{}.TableName(),
	models.User{}.TableName(),
}

func CleanDatabase() {
	gormDatabase, ok := TestDatabase.(*gorm_adapter.GormAdapter)
	if !ok {
		panic(gormDatabase)
	}

	if gormDatabase.Conn == nil {
		if err := gormDatabase.Connect(); err != nil {
			panic(err)
		}
	}

	for _, table := range tables {
		if err := gormDatabase.Conn.Exec("TRUNCATE " + table + " CASCADE").Error; err != nil {
			panic(err)
		}
	}
}
