package pet_case

import (
	core_context "pethost/app_context"
	g "pethost/frameworks/database/gateways/pet_gateway"
	"pethost/utils"
)

type CreateInput struct {
	Name      string `validate:"required"`
	Breed     string
	Birthdate string
	Gender    string
	Weight    uint8
	Species   string
	Neutered  bool
}

func (c PetCase) Create(ctx *core_context.AppContext, input *CreateInput) (string, error) {
	if err := utils.Validator.Struct(input); err != nil {
		return "", err
	}

	return c.gateway.Create(g.CreateInput{
		Name:      input.Name,
		Breed:     input.Breed,
		Birthdate: input.Birthdate,
		Gender:    input.Gender,
		Weight:    input.Weight,
		Species:   input.Species,
		UserID:    ctx.Session.UserID,
		Neutered:  input.Neutered,
	})
}
