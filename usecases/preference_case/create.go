package preference_case

import (
	"pethost/app_context"
	g "pethost/frameworks/database/gateways/preference_gateway"
	"pethost/usecases/errors_case"
	"pethost/utils"
)

type CreateInput struct {
	OnlyVaccinated          *bool  `validate:"required"`
	AcceptElderly           *bool  `validate:"required"`
	AcceptOnlyNeuteredMales *bool  `validate:"required"`
	AcceptFemales           *bool  `validate:"required"`
	DaysOfMonth             uint32 `validate:"required"`
	AcceptFemaleInHeat      *bool  `validate:"required"`
	AcceptPuppies           *bool  `validate:"required"`
	AcceptMales             *bool  `validate:"required"`
	PetWeight               uint8  `validate:"required"`
}

func (c PreferenceCase) Create(ctx *app_context.AppContext, input *CreateInput) (string, error) {
	if err := utils.Validator.Struct(input); err != nil {
		return "", errors_case.BadRequest(err)
	}

	return c.gateway.Create(g.CreateInput{
		OnlyVaccinated:          *input.OnlyVaccinated,
		AcceptElderly:           *input.AcceptElderly,
		AcceptOnlyNeuteredMales: *input.AcceptOnlyNeuteredMales,
		AcceptFemales:           *input.AcceptFemales,
		DaysOfMonth:             input.DaysOfMonth,
		AcceptFemaleInHeat:      *input.AcceptFemaleInHeat,
		AcceptPuppies:           *input.AcceptPuppies,
		AcceptMales:             *input.AcceptMales,
		PetWeight:               input.PetWeight,
		UserID:                  ctx.Session.UserID,
	})
}
