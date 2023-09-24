package pet_case

import (
	"context"
	g "pethost/adapters/gateways/pet_gateway"
)

type PatchFilter struct {
	Breed     *string
	Size      *string
	Birthdate *string
	Gender    *string
	Weight    *string
	Species   *string
	Name      *string
}

type PatchValues struct {
	Breed     string
	Size      string
	Birthdate string
	Gender    string
	Weight    string
	Species   string
	Name      string
}

func (c PetCase) Patch(ctx context.Context, filter PatchFilter, values PatchValues) (bool, error) {
	return c.gateway.Patch(
		g.PatchFilter{
			Breed:     filter.Breed,
			Size:      filter.Size,
			Birthdate: filter.Birthdate,
			Gender:    filter.Gender,
			Weight:    filter.Weight,
			Species:   filter.Species,
			Name:      filter.Name,
		}, g.PatchValues{
			Breed:     values.Breed,
			Size:      values.Size,
			Birthdate: values.Birthdate,
			Gender:    values.Gender,
			Weight:    values.Weight,
			Species:   values.Species,
			Name:      values.Name,
		})
}
