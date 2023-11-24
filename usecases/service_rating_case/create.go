package service_rating_case

import (
	"context"
	g "pethost/frameworks/database/gateways/service_rating_gateway"
)

type CreateInput struct {
	TutorID    string
	ScheduleID string
	Date       string
	Rating     int8
	Comment    string
}

func (c ServiceRatingCase) Create(ctx context.Context, input *CreateInput) (string, error) {
	return c.gateway.Create(g.CreateInput{
		TutorID:    input.TutorID,
		ScheduleID: input.ScheduleID,
		Date:       input.Date,
		Rating:     input.Rating,
		Comment:    input.Comment,
	})
}
