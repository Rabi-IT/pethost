package fixtures

import (
	"pethost/adapters/database"
	"pethost/config"
	"pethost/frameworks/database/gorm/models"
)

var TestDatabase = database.NewGorm(config.TestDatabase)

var tables = []string{
	models.Pet{}.TableName(),
}

func CleanDatabase() {
	if TestDatabase.Conn == nil {
		TestDatabase.Connect()
	}

	testConn := TestDatabase.Conn
	for _, table := range tables {
		if err := testConn.Exec("TRUNCATE " + table + " CASCADE").Error; err != nil {
			panic(err)
		}
	}
}
