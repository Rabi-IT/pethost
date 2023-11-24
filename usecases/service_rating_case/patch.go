package service_rating_case

import (
	"context"
	g "pethost/frameworks/database/gateways/service_rating_gateway"
)

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

func (c ServiceRatingCase) Patch(ctx context.Context, filter PatchFilter, values PatchValues) (bool, error) {
	return c.gateway.Patch(
		g.PatchFilter{
			ID:      filter.ID,
			Comment: filter.Comment, TutorID: filter.TutorID, ScheduleID: filter.ScheduleID, Date: filter.Date, Rating: filter.Rating,
		}, g.PatchValues{
			Comment: values.Comment, TutorID: values.TutorID, ScheduleID: values.ScheduleID, Date: values.Date, Rating: values.Rating,
		})
}
