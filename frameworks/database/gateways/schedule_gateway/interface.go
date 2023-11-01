package schedule_gateway

import "pethost/frameworks/database"

type ScheduleGateway interface {
	Create(input CreateInput) (string, error)
	Patch(filter PatchFilter, values PatchValues) (bool, error)
	Paginate(filter PaginateFilter, paginate database.PaginateInput) (*PaginateOutput, error)
}

type CreateInput struct {
	PetID   string
	Status  string
	Notes   string
	TutorID string
	HostID  string
	Date    string
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
	Status  string
	Notes   *string
	TutorID string
	HostID  string
	Date    string
	PetID   string
}

type PaginateFilter struct {
	Notes   *string
	TutorID *string
	HostID  *string
	Date    *string
	PetID   *string
	Status  *string
}

type PaginateData struct {
	Notes   *string
	TutorID string
	HostID  string
	Date    string
	PetID   string
	Status  string
}

type PaginateOutput struct {
	Data     []PaginateData
	MaxPages int
}
