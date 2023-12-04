package service_rating_case

import (
	"pethost/app_context"
	g "pethost/frameworks/database/gateways/service_rating_gateway"
	"time"
)

type CreateInput struct {
	HostID  string
	Date    string
	Rating  int8
	Comment string
}

func (c ServiceRatingCase) Create(ctx *app_context.AppContext, input *CreateInput) (string, error) {
	return c.gateway.Create(g.CreateInput{
		TutorID: ctx.Session.UserID,
		HostID:  input.HostID,
		Date:    time.Now(),
		Rating:  input.Rating,
		Comment: input.Comment,
	})
}
