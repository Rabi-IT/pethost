package host_case

import (
	"context"
	g "pethost/adapters/gateways/host_gateway"
)

func (c PetHostCase) GetByID(ctx context.Context, id string) (*g.GetByIDOutput, error) {
	return c.gateway.GetByID(id)
}
