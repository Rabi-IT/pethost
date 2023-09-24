package factories

import (
	"pethost/adapters/database"
	"pethost/config"
)

func NewProductionDatabase() database.Database {
	return database.NewGorm(config.ProductionDatabase)
}
