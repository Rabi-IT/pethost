package pet_case

import (
	"pethost/app_context"
	g "pethost/frameworks/database/gateways/pet_gateway"
)

type ListInput struct {
	PetIDs []string
}

func (c PetCase) List(ctx *app_context.AppContext, input *ListInput) ([]g.ListOutput, error) {
	filter := g.ListInput{
		PetIDs: input.PetIDs,
	}

	if ctx.Session.Role.IsUser() {
		filter.TutorID = &ctx.Session.UserID
	}

	return c.gateway.List(filter)
}
