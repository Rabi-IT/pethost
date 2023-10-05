package pethost_case

import (
	"context"
	g "pethost/adapters/gateways/pethost_gateway"
)

type PatchFilter struct {
	ID string
}

type PatchValues struct {
	ZIP            string
	Phone          string
	Email          string
	EmergencyPhone string
	Neighborhood   string
	Street         string
	SocialID       string
	TaxID          string
	City           string
	State          string
	Complement     string
	Name           string
}

func (c PetHostCase) Patch(ctx context.Context, filter PatchFilter, values PatchValues) (bool, error) {
	return c.gateway.Patch(g.PatchFilter{
		ID: filter.ID,
	}, g.PatchValues{
		ZIP:            values.ZIP,
		Phone:          values.Phone,
		Email:          values.Email,
		EmergencyPhone: values.EmergencyPhone,
		Neighborhood:   values.Neighborhood,
		Street:         values.Street,
		City:           values.City,
		State:          values.State,
		Complement:     values.Complement,
		Name:           values.Name,
	})
}
