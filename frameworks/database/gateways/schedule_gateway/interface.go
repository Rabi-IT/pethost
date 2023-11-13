package schedule_gateway

import (
	"pethost/frameworks/database"
	"pethost/usecases/schedule_case/schedule_status"
	"time"
)

type ScheduleGateway interface {
	Create(input CreateInput) (string, error)
	Patch(filter PatchFilter, values PatchValues) (bool, error)
	Paginate(filter PaginateFilter, paginate database.PaginateInput) (*PaginateOutput, error)
}

type CreateInput struct {
	PetID       string
	TutorID     string
	HostID      string
	MonthYear   time.Time
	DaysOfMonth uint32
	Status      schedule_status.Status
	Notes       string
}

type PatchFilter struct {
	ID      *string
	Status  *string
	Notes   *string
	TutorID *string
	HostID  *string
	Date    *string
	PetID   *string
}

type PatchValues struct {
	Status string
}

type PaginateFilter struct {
	HostID string
	Status schedule_status.Status
}

type PaginateData struct {
	PetID       string
	TutorID     string
	MonthYear   time.Time
	DaysOfMonth uint32
	Status      schedule_status.Status
	Notes       string
}

type PaginateOutput struct {
	Data     []PaginateData
	MaxPages int
}
