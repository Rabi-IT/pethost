package user_case

import (
	"context"
	g "pethost/frameworks/database/gateways/user_gateway"
	"pethost/utils"
)

type CreateInput struct {
	Name           string `validate:"required"`
	Photo          string
	TaxID          string `validate:"required"`
	City           string `validate:"required"`
	State          string `validate:"required"`
	Phone          string `validate:"required"`
	ZIP            string `validate:"required"`
	SocialID       string `validate:"required"`
	Email          string `validate:"required"`
	EmergencyPhone string
	Neighborhood   string `validate:"required"`
	Street         string `validate:"required"`
	Complement     string
}

func (c UserCase) Create(ctx context.Context, input *CreateInput) (string, error) {
	if err := utils.Validator.Struct(input); err != nil {
		return "", err
	}

	return c.gateway.Create(g.CreateInput{
		City:           input.City,
		State:          input.State,
		ZIP:            input.ZIP,
		Phone:          input.Phone,
		Email:          input.Email,
		Photo:          input.Photo,
		TaxID:          input.TaxID,
		SocialID:       input.SocialID,
		Street:         input.Street,
		Complement:     input.Complement,
		EmergencyPhone: input.EmergencyPhone,
		Neighborhood:   input.Neighborhood,
		Name:           input.Name,
	})
}
