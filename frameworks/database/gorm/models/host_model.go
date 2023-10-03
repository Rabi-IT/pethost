package models

type PetHost struct {
	ID             string `gorm:"type:uuid"`
	Name           string `gorm:"not null"`
	TaxID          string `gorm:"not null"`
	City           string `gorm:"not null"`
	State          string `gorm:"not null"`
	Complement     string
	Phone          string `gorm:"not null"`
	ZIP            string `gorm:"not null"`
	SocialID       string `gorm:"not null"`
	Email          string `gorm:"not null"`
	EmergencyPhone string `gorm:"not null"`
	Neighborhood   string `gorm:"not null"`
	Street         string `gorm:"not null"`
}

func (m PetHost) TableName() string {
	return "pet_hosts"
}
