package host_case

import (
	"context"
	g "pethost/adapters/gateways/host_gateway"
)

type CreateInput struct {
	Name           string `validate:"required"`
	TaxID          string `validate:"required"`
	City           string `validate:"required"`
	State          string `validate:"required"`
	Phone          string `validate:"required"`
	ZIP            string `validate:"required"`
	SocialID       string `validate:"required"`
	Email          string `validate:"required"`
	EmergencyPhone string `validate:"required"`
	Neighborhood   string `validate:"required"`
	Street         string `validate:"required"`
	Complement     string
}

func (c PetHostCase) Create(ctx context.Context, input *CreateInput) (string, error) {
	return c.gateway.Create(g.CreateInput{
		Name:           input.Name,
		TaxID:          input.TaxID,
		City:           input.City,
		State:          input.State,
		Complement:     input.Complement,
		Phone:          input.Phone,
		ZIP:            input.ZIP,
		SocialID:       input.SocialID,
		Email:          input.Email,
		EmergencyPhone: input.EmergencyPhone,
		Neighborhood:   input.Neighborhood,
		Street:         input.Street,
	})
}
