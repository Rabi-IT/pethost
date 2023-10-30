package preference_gateway

type PreferenceGateway interface {
	Create(input CreateInput) (string, error)
}

type CreateInput struct {
	AcceptPuppies           bool
	AcceptMales             bool
	PetWeight               uint8
	AcceptFemaleInHeat      bool
	AcceptElderly           bool
	AcceptOnlyNeuteredMales bool
	AcceptFemales           bool
	DaysOfMonth             uint32
	OnlyVaccinated          bool
}
