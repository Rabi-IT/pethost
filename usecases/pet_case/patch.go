package pet_case

import (
	"context"
	g "pethost/frameworks/database/gateways/pet_gateway"
	"pethost/usecases/pet_case/pet"
	"time"
)

type PatchFilter struct {
	ID      *string
	Breed   *string
	Gender  *pet.Gender
	Weight  *uint8
	Species *string
	Name    *string
}

type PatchValues struct {
	Breed      string
	Birthdate  time.Time
	Gender     pet.Gender
	Weight     uint8
	Species    string
	Name       string
	Neutered   *bool
	Vaccinated *bool
}

func (c PetCase) Patch(ctx context.Context, filter PatchFilter, values PatchValues) (bool, error) {
	return c.gateway.Patch(
		g.PatchFilter{
			ID:      filter.ID,
			Breed:   filter.Breed,
			Gender:  filter.Gender,
			Weight:  filter.Weight,
			Species: filter.Species,
			Name:    filter.Name,
		}, g.PatchValues{
			Breed:      values.Breed,
			Birthdate:  values.Birthdate,
			Gender:     values.Gender,
			Weight:     values.Weight,
			Species:    values.Species,
			Name:       values.Name,
			Neutered:   values.Neutered,
			Vaccinated: values.Vaccinated,
		})
}
