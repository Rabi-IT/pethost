package tutor_case

import (
	"context"
	g "pethost/adapters/gateways/tutor_gateway"
)

func (c TutorCase) GetByID(ctx context.Context, id string) (*g.GetByIDOutput, error) {
	return c.gateway.GetByID(id)
}
