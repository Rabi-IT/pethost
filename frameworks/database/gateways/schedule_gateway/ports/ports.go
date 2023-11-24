package ports

import (
	"pethost/usecases/schedule_case/schedule"
	"pethost/usecases/schedule_case/schedule_status"
	"time"
)

type ScheduleDate struct {
	MonthYear   time.Time            `validate:"required" json:"monthYear"`
	DaysOfMonth schedule.DaysOfMonth `validate:"required" json:"daysOfMonth"`
}

type ScheduleHistory struct {
	UserID    string                 `json:"userId"`
	NewStatus schedule_status.Status `json:"newStatus"`
	Date      time.Time              `json:"date"`
	Notes     string                 `json:"notes"`
}
