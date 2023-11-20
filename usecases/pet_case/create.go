package pet_case

import (
	"errors"
	"fmt"
	core_context "pethost/app_context"
	g "pethost/frameworks/database/gateways/pet_gateway"
	"pethost/usecases/pet_case/pet"
	"pethost/utils"
	"time"
)

type CreateInput struct {
	Name       string `validate:"required"`
	Breed      string
	Birthdate  time.Time  `validate:"required"`
	Gender     pet.Gender `validate:"required"`
	Weight     pet.Weight `validate:"required"`
	Species    pet.Specie `validate:"required"`
	Neutered   *bool      `validate:"required"`
	Vaccinated *bool      `validate:"required"`
}

var ErrInvalidWeight = errors.New("invalid weight")

func (c PetCase) Create(ctx *core_context.AppContext, input *CreateInput) (string, error) {
	if err := utils.Validator.Struct(input); err != nil {
		return "", err
	}

	if input.Weight&(input.Weight-1) != 0 {
		return "", fmt.Errorf("Create %w: %d", ErrInvalidWeight, input.Weight)
	}

	return c.gateway.Create(g.CreateInput{
		Name:       input.Name,
		Breed:      input.Breed,
		Birthdate:  input.Birthdate,
		Gender:     input.Gender,
		Weight:     input.Weight,
		Species:    input.Species,
		UserID:     ctx.Session.UserID,
		Neutered:   *input.Neutered,
		Vaccinated: *input.Vaccinated,
	})
}
