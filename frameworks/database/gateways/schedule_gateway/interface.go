package schedule_gateway

import (
	"pethost/frameworks/database"
	"pethost/frameworks/database/gateways/schedule_gateway/ports"
	"pethost/usecases/schedule_case/schedule_status"
	"time"
)

type ScheduleGateway interface {
	Create(input CreateInput) (string, error)
	Patch(filter PatchFilter, values PatchValues) (bool, error)
	Paginate(filter PaginateFilter, paginate database.PaginateInput) (*PaginateOutput, error)
	GetByID(id string) (*GetByIDOutput, error)
}

type CreateInput struct {
	TutorID   string
	HostID    string
	Status    schedule_status.Status
	Notes     string
	PetIDs    []string
	StartDate time.Time
	EndDate   time.Time
}

type PatchFilter struct {
	ID       string
	Status   *schedule_status.Status
	StatusOR []schedule_status.Status
	TutorID  *string
	HostID   *string
}

type PatchValues struct {
	Status  *schedule_status.Status
	History *ports.ScheduleHistory
}

type PaginateFilter struct {
	TutorID *string
	HostID  *string
	Status  *schedule_status.Status
}

type PaginateData struct {
	PetIDs    []string
	TutorID   string
	StartDate time.Time
	EndDate   time.Time
	Status    schedule_status.Status
	Notes     string
}

type PaginateOutput struct {
	Data     []PaginateData
	MaxPages int
}

type GetByIDOutput struct {
	Status schedule_status.Status
}
