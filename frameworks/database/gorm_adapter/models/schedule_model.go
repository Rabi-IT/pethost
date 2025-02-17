package models

import (
	"pethost/frameworks/database/gateways/schedule_gateway/ports"
	"pethost/usecases/schedule_case/schedule_status"
	"time"

	"gorm.io/datatypes"
)

type Schedule struct {
	ID        string `gorm:"type:uuid"`
	TutorID   string `gorm:"not null; uniqueIndex:unique_schedule_date"`
	HostID    string `gorm:"not null; uniqueIndex:unique_schedule_date"`
	PetIDs    datatypes.JSONSlice[string]
	StartDate time.Time
	EndDate   time.Time
	Tutor     User
	Host      User

	Status schedule_status.Status `gorm:"not null"`
	Notes  string

	History   datatypes.JSONSlice[ports.ScheduleHistory] `gorm:"default:'[]'::jsonb"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (m Schedule) TableName() string {
	return "schedules"
}
