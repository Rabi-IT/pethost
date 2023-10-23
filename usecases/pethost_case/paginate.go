package pethost_case

import (
	"context"
	"pethost/adapters/database"
	g "pethost/adapters/gateways/pethost_gateway"
)

type PaginateFilter struct {
	Name       *string
	City       *string
	State      *string
	Complement *string

	ZIP          *string
	Email        *string
	Neighborhood *string
}

func (c PetHostCase) Paginate(ctx context.Context, input PaginateFilter, paginate database.PaginateInput) (*g.PaginateOutput, error) {
	return c.gateway.Paginate(g.PaginateFilter{
		Name:         input.Name,
		City:         input.City,
		State:        input.State,
		Neighborhood: input.Neighborhood,
	}, paginate)
}
