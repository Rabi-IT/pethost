package pet_case

import (
	"pethost/app_context"
	g "pethost/frameworks/database/gateways/pet_gateway"
)

func (c PetCase) GetByID(ctx *app_context.AppContext, id string) (*g.GetByIDOutput, error) {
	return c.gateway.GetByID(id)
}
