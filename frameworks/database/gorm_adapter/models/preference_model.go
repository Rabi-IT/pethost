package models

type Preference struct {
	ID                      string `gorm:"type:uuid"`
	DaysOfMonth             uint32 `gorm:"not null"`
	OnlyVaccinated          bool   `gorm:"not null"`
	AcceptElderly           bool   `gorm:"not null"`
	AcceptOnlyNeuteredMales bool   `gorm:"not null"`
	AcceptFemales           bool   `gorm:"not null"`
	PetWeight               uint8  `gorm:"not null"`
	AcceptFemaleInHeat      bool   `gorm:"not null"`
	AcceptPuppies           bool   `gorm:"not null"`
	AcceptMales             bool   `gorm:"not null"`
	PetCapacity             uint8  `gorm:"not null"`

	User   User
	UserID string `gorm:"not null"`
}

func (m Preference) TableName() string {
	return "preferences"
}
