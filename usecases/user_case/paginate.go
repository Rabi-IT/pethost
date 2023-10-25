package user_case

import (
	"context"
	"pethost/adapters/database"
	g "pethost/adapters/gateways/user_gateway"
)

type PaginateFilter struct {
	City           *string
	State          *string
	ZIP            *string
	Phone          *string
	Email          *string
	Photo          *string
	TaxID          *string
	SocialID       *string
	Street         *string
	Complement     *string
	EmergencyPhone *string
	Name           *string
}

func (c UserCase) Paginate(ctx context.Context, input PaginateFilter, paginate database.PaginateInput) (*g.PaginateOutput, error) {
	return c.gateway.Paginate(g.PaginateFilter{
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
		Name:           input.Name,
	}, paginate)
}
