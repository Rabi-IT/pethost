package models

import (
	"pethost/usecases/schedule_case/schedule_status"
	"time"
)

type Schedule struct {
	ID          string    `gorm:"type:uuid"`
	TutorID     string    `gorm:"not null; uniqueIndex:unique_schedule_date"`
	HostID      string    `gorm:"not null; uniqueIndex:unique_schedule_date"`
	PetID       string    `gorm:"not null; uniqueIndex:unique_schedule_date"`
	DaysOfMonth uint32    `gorm:"not null; uniqueIndex:unique_schedule_date"`
	MonthYear   time.Time `gorm:"not null; uniqueIndex:unique_schedule_date"`

	Tutor User
	Host  User
	Pet   Pet

	Status    schedule_status.Status `gorm:"not null"`
	Notes     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (m Schedule) TableName() string {
	return "schedules"
}
