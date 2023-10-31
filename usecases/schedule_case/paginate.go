package schedule_case

import (
	"context"
	"pethost/frameworks/database"
	g "pethost/frameworks/database/gateways/schedule_gateway"
)

type PaginateFilter struct {
	Date    *string
	PetID   *string
	Status  *string
	Notes   *string
	TutorID *string
	HostID  *string
}

func (c ScheduleCase) Paginate(ctx context.Context, input PaginateFilter, paginate database.PaginateInput) (*g.PaginateOutput, error) {
	return c.gateway.Paginate(g.PaginateFilter{
		Date:    input.Date,
		PetID:   input.PetID,
		Status:  input.Status,
		Notes:   input.Notes,
		TutorID: input.TutorID,
		HostID:  input.HostID,
	}, paginate)
}
