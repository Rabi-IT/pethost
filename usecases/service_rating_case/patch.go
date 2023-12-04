package service_rating_case

import (
	"pethost/app_context"
	g "pethost/frameworks/database/gateways/service_rating_gateway"
)

type PatchFilter struct {
	ID *string
}

type PatchValues struct {
	Comment string
	Rating  int8
}

func (c ServiceRatingCase) Patch(ctx *app_context.AppContext, filter PatchFilter, values PatchValues) (bool, error) {
	return c.gateway.Patch(
		g.PatchFilter{
			ID:      filter.ID,
			TutorID: &ctx.Session.UserID,
		}, g.PatchValues{
			Comment: values.Comment,
			Rating:  values.Rating,
		})
}
