package database

import (
	"pethost/frameworks/database/gorm/models"
)

func (d *GormAdapter) Migrate() error {
	return d.Conn.AutoMigrate(
		&models.Pet{},
		&models.PetHost{},
		&models.Tutor{},
	)
}
