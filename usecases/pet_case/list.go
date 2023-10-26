package pet_case

import (
	"context"
	g "pethost/frameworks/database/gateways/pet_gateway"
)

type ListInput struct {
	Weight    *string
	Species   *string
	Name      *string
	Breed     *string
	Size      *string
	Birthdate *string
	Gender    *string
}

func (c PetCase) List(ctx context.Context, input ListInput) ([]g.ListOutput, error) {
	return c.gateway.List(g.ListInput{
		Weight:    input.Weight,
		Species:   input.Species,
		Name:      input.Name,
		Breed:     input.Breed,
		Size:      input.Size,
		Birthdate: input.Birthdate,
		Gender:    input.Gender,
	})
}
