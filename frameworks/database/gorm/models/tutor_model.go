package models

type Tutor struct {
	ID             string `gorm:"type:uuid"`
	SocialID       string
	Street         string
	Complement     string
	EmergencyPhone string
	Name           string `gorm:"not null"`
	Email          string `gorm:"not null"`
	Photo          string
	TaxID          string `gorm:"not null"`
	Phone          string `gorm:"not null"`
	City           string
	State          string
	ZIP            string
	Neighborhood   string
}

func (m Tutor) TableName() string {
	return "tutors"
}
