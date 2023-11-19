package ports

import "time"

type ScheduleDate struct {
	MonthYear   time.Time `validate:"required" json:"monthYear"`
	DaysOfMonth uint32    `validate:"required" json:"daysOfMonth"`
}
