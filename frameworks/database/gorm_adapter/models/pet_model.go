package models

import (
	"pethost/usecases/pet_case/pet"
	"time"
)

type Pet struct {
	ID         string     `gorm:"type:uuid"`
	Weight     pet.Weight `gorm:"not null"`
	Species    pet.Specie `gorm:"not null"`
	Name       string     `gorm:"not null"`
	Breed      string     `gorm:"not null"`
	Photo      string     `gorm:"not null"`
	Birthdate  time.Time  `gorm:"not null"`
	Gender     pet.Gender `gorm:"not null"`
	Neutered   bool       `gorm:"not null"`
	Vaccinated bool       `gorm:"not null"`
	Tutor      User
	TutorID    string `gorm:"not null"`
}

func (m Pet) TableName() string {
	return "pets"
}
