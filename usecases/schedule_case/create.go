package schedule_case

import (
	"pethost/app_context"
	"pethost/frameworks/database/gateways/pet_gateway"
	"pethost/frameworks/database/gateways/preference_gateway"
	g "pethost/frameworks/database/gateways/schedule_gateway"
	"pethost/usecases/pet_case"
	"pethost/usecases/pet_case/pet"
	"pethost/usecases/schedule_case/schedule"
	"pethost/usecases/schedule_case/schedule_status"
	"pethost/utils"
	"time"
)

type CreateInput struct {
	HostID        string    `validate:"required"`
	PetIDs        []string  `validate:"required"`
	StartDate     time.Time `validate:"required"`
	EndDate       time.Time `validate:"required"`
	FemalesInHeat map[string]bool
	Notes         string
}

func (c *ScheduleCase) Create(ctx *app_context.AppContext, input *CreateInput) (id string, err error) {
	if err := utils.Validator.Struct(input); err != nil {
		return "", err
	}

	petsFound, err := c.pet.List(ctx, &pet_case.ListInput{
		PetIDs: input.PetIDs,
	})

	if err != nil || len(petsFound) == 0 {
		return
	}

	preference, err := c.preference.GetByFilter(ctx, &preference_gateway.GetByFilterInput{
		UserID: input.HostID,
	})

	if err != nil || preference == nil {
		return
	}

	if valid := c.validateSchedule(preference, input.StartDate, input.EndDate, petsFound, input.FemalesInHeat); !valid {
		return
	}

	return c.gateway.Create(g.CreateInput{
		PetIDs:    input.PetIDs,
		TutorID:   ctx.Session.UserID,
		HostID:    input.HostID,
		StartDate: input.StartDate,
		EndDate:   input.EndDate,
		Status:    schedule_status.Open,
		Notes:     input.Notes,
	})
}

func (*ScheduleCase) validateSchedule(
	preference *preference_gateway.GetByFilterOutput,
	startDate time.Time,
	endDate time.Time,
	pets []pet_gateway.ListOutput,
	femalesInHeat map[string]bool,
) bool {
	if femalesInHeat == nil {
		femalesInHeat = make(map[string]bool)
	}

	if startDate.Before(time.Now()) {
		return false
	}

	if startDate.After(endDate) {
		return false
	}

	dates := schedule.ToSchedule(startDate, endDate)

	for _, d := range dates {
		if (preference.DaysOfMonth & d.DaysOfMonth) != d.DaysOfMonth {
			return false
		}
	}

	for _, p := range pets {
		if pet.Female == p.Gender && !isValidFemale(preference, p, femalesInHeat[p.ID]) {
			return false
		}

		if !isValidMale(preference, p) {
			return false
		}

		if preference.PetWeight&p.Weight == 0 {
			return false
		}

		petYears := pet.CalculateAge(p.Birthdate, time.Now())
		if !preference.AcceptPuppies && petYears < pet.PuppieAge {
			return false
		} else if !preference.AcceptElderly && petYears >= pet.ElderlyAge {
			return false
		}

		if preference.OnlyVaccinated && !p.Vaccinated {
			return false
		}
	}

	return true
}

func isValidFemale(preference *preference_gateway.GetByFilterOutput, pet pet_gateway.ListOutput, femaleInHeat bool) bool {
	if !preference.AcceptFemales {
		return false
	}

	if !femaleInHeat || pet.Neutered {
		return true
	}

	return preference.AcceptFemaleInHeat
}

func isValidMale(preference *preference_gateway.GetByFilterOutput, pet pet_gateway.ListOutput) bool {
	if !preference.AcceptMales {
		return false
	}

	if !preference.AcceptOnlyNeuteredMales {
		return true
	}

	return pet.Neutered
}
