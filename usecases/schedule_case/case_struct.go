package schedule_case

import (
	g "pethost/frameworks/database/gateways/schedule_gateway"
	"pethost/usecases/pet_case"
	"pethost/usecases/preference_case"
)

type ScheduleCase struct {
	gateway    g.ScheduleGateway
	preference *preference_case.PreferenceCase
	pet        *pet_case.PetCase
}

func New(
	gateway g.ScheduleGateway,
	preference *preference_case.PreferenceCase,
	pet *pet_case.PetCase,
) *ScheduleCase {
	return &ScheduleCase{gateway, preference, pet}
}
