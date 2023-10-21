package tutor_case

import (
	"context"
	g "pethost/adapters/gateways/tutor_gateway"
)

type ListInput struct {
	EmergencyPhone *string
	Name           *string
	Email          *string
	Photo          *string
	TaxID          *string
	SocialID       *string
	Street         *string
	Complement     *string
	Phone          *string
	City           *string
	State          *string
	ZIP            *string
}

func (c TutorCase) List(ctx context.Context, input ListInput) ([]g.ListOutput, error) {
	return c.gateway.List(g.ListInput{
		EmergencyPhone: input.EmergencyPhone,
		Name:           input.Name,
		Email:          input.Email,
		Photo:          input.Photo,
		TaxID:          input.TaxID,
		SocialID:       input.SocialID,
		Street:         input.Street,
		Complement:     input.Complement,
		Phone:          input.Phone,
		City:           input.City,
		State:          input.State,
		ZIP:            input.ZIP,
	})
}
