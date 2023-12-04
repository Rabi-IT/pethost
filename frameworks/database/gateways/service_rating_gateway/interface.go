package service_rating_gateway

import (
	"pethost/frameworks/database"
	"time"
)

type ServiceRatingGateway interface {
	Create(input CreateInput) (string, error)
	Paginate(filter PaginateFilter, paginate database.PaginateInput) (*PaginateOutput, error)
	Patch(filter PatchFilter, values PatchValues) (bool, error)
	Delete(id string) (bool, error)
}

type CreateInput struct {
	TutorID string
	HostID  string
	Date    time.Time
	Rating  int8
	Comment string
}

type PaginateFilter struct {
	HostID *string
	Date   *time.Time
	Rating *int8
}

type PaginateData struct {
	TutorID string
	HostID  string
	Date    time.Time
	Rating  int8
	Comment string
}

type PaginateOutput struct {
	Data     []PaginateData
	MaxPages int
}

type PatchFilter struct {
	ID      *string
	TutorID *string
}

type PatchValues struct {
	Comment string
	Rating  int8
}

type DeleteInput struct {
	ID      string
	TutorID string
}
