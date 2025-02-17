package models

import (
	"pethost/usecases/pet_case/pet"
	"pethost/usecases/schedule_case/schedule"
)

type Preference struct {
	ID                      string               `gorm:"type:uuid"`
	DaysOfMonth             schedule.DaysOfMonth `gorm:"not null"`
	OnlyVaccinated          bool                 `gorm:"not null"`
	AcceptElderly           bool                 `gorm:"not null"`
	AcceptOnlyNeuteredMales bool                 `gorm:"not null"`
	AcceptFemales           bool                 `gorm:"not null"`
	PetWeight               pet.Weight           `gorm:"not null"`
	AcceptFemaleInHeat      bool                 `gorm:"not null"`
	AcceptPuppies           bool                 `gorm:"not null"`
	AcceptMales             bool                 `gorm:"not null"`
	PetCapacity             uint8                `gorm:"not null"`

	User   User
	UserID string `gorm:"not null"`
}

func (m Preference) TableName() string {
	return "preferences"
}
