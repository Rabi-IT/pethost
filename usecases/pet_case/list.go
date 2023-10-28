package pet_case

import (
	"pethost/app_context"
	g "pethost/frameworks/database/gateways/pet_gateway"
)

func (c PetCase) List(ctx *app_context.AppContext) ([]g.ListOutput, error) {
	filter := g.ListInput{}

	if ctx.Session.Role.IsUser() {
		filter.TutorID = &ctx.Session.UserID
	}

	return c.gateway.List(filter)
}
