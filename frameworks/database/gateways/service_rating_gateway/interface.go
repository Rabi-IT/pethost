package service_rating_gateway

import (
	"pethost/frameworks/database"
)

type ServiceRatingGateway interface {
	Create(input CreateInput) (string, error)
	List(input ListInput) ([]ListOutput, error)
	Paginate(filter PaginateFilter, paginate database.PaginateInput) (*PaginateOutput, error)
	Patch(filter PatchFilter, values PatchValues) (bool, error)
	Delete(id string) (bool, error)
}

type CreateInput struct {
	TutorID    string
	ScheduleID string
	Date       string
	Rating     int8
	Comment    string
}

type ListInput struct {
	Date       *string
	Rating     *int8
	Comment    *string
	TutorID    *string
	ScheduleID *string
}

type ListOutput struct {
	Date       string
	Rating     int8
	Comment    string
	TutorID    string
	ScheduleID string
}

type PaginateFilter struct {
	TutorID    *string
	ScheduleID *string
	Date       *string
	Rating     *int8
	Comment    *string
}

type PaginateData struct {
	TutorID    string
	ScheduleID string
	Date       string
	Rating     int8
	Comment    string
}

type PaginateOutput struct {
	Data     []PaginateData
	MaxPages int
}

type PatchFilter struct {
	ID         *string
	Comment    *string
	TutorID    *string
	ScheduleID *string
	Date       *string
	Rating     *int8
}

type PatchValues struct {
	Comment    string
	TutorID    string
	ScheduleID string
	Date       string
	Rating     int8
}

type DeleteInput struct {
	TutorID    string
	ScheduleID string
	Date       string
	Rating     int8
	Comment    string
}
