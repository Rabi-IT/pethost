package service_rating_case

import (
	"context"
	"pethost/frameworks/database"
	g "pethost/frameworks/database/gateways/service_rating_gateway"
	"time"
)

type PaginateFilter struct {
	Date   *time.Time
	Rating *int8
	HostID *string
}

func (c ServiceRatingCase) Paginate(ctx context.Context, input PaginateFilter, paginate database.PaginateInput) (*g.PaginateOutput, error) {
	return c.gateway.Paginate(g.PaginateFilter{
		HostID: input.HostID,
		Date:   input.Date,
		Rating: input.Rating,
	}, paginate)
}
