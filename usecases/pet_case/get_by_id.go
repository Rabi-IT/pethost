package pet_case

import (
	"context"
	g "pethost/adapters/gateways/pet_gateway"
)

func (c PetCase) GetByID(ctx context.Context, id string) (*g.GetByIDOutput, error) {
	return c.gateway.GetByID(id)
}
