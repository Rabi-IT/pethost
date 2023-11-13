package schedule_case

import (
	"pethost/app_context"
	"pethost/frameworks/database/gateways/preference_gateway"
	g "pethost/frameworks/database/gateways/schedule_gateway"
	"pethost/usecases/schedule_case/schedule_status"
	"pethost/utils"
	"time"
)

type CreateInput struct {
	PetID        string    `validate:"required"`
	HostID       string    `validate:"required"`
	MonthYear    time.Time `validate:"required"`
	DaysOfMonth  uint32    `validate:"required"`
	Notes        string
	FemaleInHeat *bool
}

func (c *ScheduleCase) Create(ctx *app_context.AppContext, input *CreateInput) (id string, err error) {
	if err := utils.Validator.Struct(input); err != nil {
		return "", err
	}

	// verificar UserID na hora de buscar o pet
	pet, err := c.pet.GetByID(ctx, input.PetID)
	if err != nil {
		return
	}

	filter := &preference_gateway.GetByFilterInput{
		UserID:         input.HostID,
		AcceptPuppies:  nil,
		PetWeight:      pet.Weight,
		AcceptElderly:  nil,
		DaysOfMonth:    input.DaysOfMonth,
		OnlyVaccinated: nil,
	}

	isFemale := pet.Gender == "female"
	True := true
	False := false
	if isFemale {
		filter.AcceptFemales = &True
		filter.AcceptFemaleInHeat = input.FemaleInHeat
	} else {
		filter.AcceptMales = &True
		if !pet.Neutered {
			filter.AcceptOnlyNeuteredMales = &False
		}
	}

	host, err := c.preference.GetByFilter(ctx, filter)

	if err != nil || host == nil {
		return
	}

	return c.gateway.Create(g.CreateInput{
		PetID:       input.PetID,
		TutorID:     ctx.Session.UserID,
		HostID:      input.HostID,
		MonthYear:   input.MonthYear,
		DaysOfMonth: input.DaysOfMonth,
		Status:      schedule_status.Open,
		Notes:       input.Notes,
	})
}
