package preference_gateway

import "pethost/usecases/pet_case/pet"

type PreferenceGateway interface {
	Create(input CreateInput) (string, error)
	GetByFilter(filter *GetByFilterInput) (*GetByFilterOutput, error)
}

type CreateInput struct {
	UserID                  string
	AcceptPuppies           bool
	AcceptMales             bool
	PetWeight               pet.Weight
	AcceptFemaleInHeat      bool
	AcceptElderly           bool
	AcceptOnlyNeuteredMales bool
	AcceptFemales           bool
	DaysOfMonth             uint32
	OnlyVaccinated          bool
	Capacity                uint8
}

type GetByFilterInput struct {
	UserID string
}

type GetByFilterOutput struct {
	UserID                  string
	DaysOfMonth             uint32
	OnlyVaccinated          bool
	AcceptElderly           bool
	AcceptOnlyNeuteredMales bool
	AcceptFemales           bool
	PetWeight               pet.Weight
	AcceptFemaleInHeat      bool
	AcceptPuppies           bool
	AcceptMales             bool
}
