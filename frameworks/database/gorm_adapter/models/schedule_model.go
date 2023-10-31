package models

type Schedule struct {
	ID      string `gorm:"type:uuid"`
	TutorID string `gorm:"not null"`
	HostID  string `gorm:"not null"`
	Date    string `gorm:"not null"`
	PetID   string `gorm:"not null"`
	Status  string `gorm:"not null"`
	Notes   string
}

func (m Schedule) TableName() string {
	return "schedules"
}
