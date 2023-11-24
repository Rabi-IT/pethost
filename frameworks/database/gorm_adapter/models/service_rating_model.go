package models

type ServiceRating struct {
	ID         string `gorm:"type:uuid"`
	Date       string `gorm:"not null"`
	Rating     int8   `gorm:"not null"`
	Comment    string `gorm:"not null"`
	TutorID    string `gorm:"not null"`
	ScheduleID string `gorm:"not null"`
}

func (m ServiceRating) TableName() string {
	return "service_ratings"
}
