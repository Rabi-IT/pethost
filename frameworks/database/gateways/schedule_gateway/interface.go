package schedule_gateway

import (
	"pethost/frameworks/database"
	"pethost/frameworks/database/gateways/schedule_gateway/ports"
	"pethost/usecases/schedule_case/schedule_status"
)

type ScheduleGateway interface {
	Create(input CreateInput) (string, error)
	Patch(filter PatchFilter, values PatchValues) (bool, error)
	Paginate(filter PaginateFilter, paginate database.PaginateInput) (*PaginateOutput, error)
}

type CreateInput struct {
	TutorID string
	HostID  string
	Status  schedule_status.Status
	Notes   string
	PetIDs  []string
	Dates   []ports.ScheduleDate
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
	PetIDs  []string
	TutorID string
	Dates   []ports.ScheduleDate
	Status  schedule_status.Status
	Notes   string
}

type PaginateOutput struct {
	Data     []PaginateData
	MaxPages int
}
