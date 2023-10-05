package pethost_case

import (
	"context"
	g "pethost/adapters/gateways/pethost_gateway"
)

func (c PetHostCase) GetByID(ctx context.Context, id string) (*g.GetByIDOutput, error) {
	return c.gateway.GetByID(id)
}
