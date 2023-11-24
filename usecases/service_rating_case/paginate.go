package service_rating_case

import (
	"context"
	"pethost/frameworks/database"
	g "pethost/frameworks/database/gateways/service_rating_gateway"
)

type PaginateFilter struct {
	ScheduleID *string
	Date       *string
	Rating     *int8
	Comment    *string
	TutorID    *string
}

func (c ServiceRatingCase) Paginate(ctx context.Context, input PaginateFilter, paginate database.PaginateInput) (*g.PaginateOutput, error) {
	return c.gateway.Paginate(g.PaginateFilter{
		ScheduleID: input.ScheduleID,
		Date:       input.Date,
		Rating:     input.Rating,
		Comment:    input.Comment,
		TutorID:    input.TutorID,
	}, paginate)
}
