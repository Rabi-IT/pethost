package models

import "pethost/usecases/pet_case/pet"

type Pet struct {
	ID         string     `gorm:"type:uuid"`
	Weight     uint8      `gorm:"not null"`
	Species    string     `gorm:"not null"`
	Name       string     `gorm:"not null"`
	Breed      string     `gorm:"not null"`
	Birthdate  string     `gorm:"not null"`
	Gender     pet.Gender `gorm:"not null"`
	Neutered   bool       `gorm:"not null"`
	Vaccinated bool       `gorm:"not null"`
	Tutor      User
	TutorID    string `gorm:"not null"`
}

func (m Pet) TableName() string {
	return "pets"
}
