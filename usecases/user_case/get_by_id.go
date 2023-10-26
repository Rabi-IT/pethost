package user_case

import (
	"context"
	g "pethost/frameworks/database/gateways/user_gateway"
)

func (c UserCase) GetByID(ctx context.Context, id string) (*g.GetByIDOutput, error) {
	return c.gateway.GetByID(id)
}
