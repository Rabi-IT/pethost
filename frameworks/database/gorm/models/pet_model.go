package models

type Pet struct {
	ID        string `gorm:"type:uuid"`
	Weight    string `gorm:"not null"`
	Species   string `gorm:"not null"`
	Name      string `gorm:"not null"`
	Breed     string `gorm:"not null"`
	Size      string `gorm:"not null"`
	Birthdate string `gorm:"not null"`
	Gender    string `gorm:"not null"`
	Tutor     Tutor
	TutorID   string `gorm:"not null"`
}

func (m Pet) TableName() string {
	return "pets"
}
