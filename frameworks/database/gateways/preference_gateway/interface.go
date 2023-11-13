package preference_gateway

type PreferenceGateway interface {
	Create(input CreateInput) (string, error)
	GetByFilter(filter *GetByFilterInput) (*GetByFilterOutput, error)
}

type CreateInput struct {
	UserID                  string
	AcceptPuppies           bool
	AcceptMales             bool
	PetWeight               uint8
	AcceptFemaleInHeat      bool
	AcceptElderly           bool
	AcceptOnlyNeuteredMales bool
	AcceptFemales           bool
	DaysOfMonth             uint32
	OnlyVaccinated          bool
	Capacity                uint8
}

type GetByFilterInput struct {
	UserID                  string
	AcceptPuppies           *bool
	AcceptMales             *bool
	PetWeight               uint8
	AcceptFemaleInHeat      *bool
	AcceptElderly           *bool
	AcceptOnlyNeuteredMales *bool
	AcceptFemales           *bool
	DaysOfMonth             uint32
	OnlyVaccinated          *bool
}

type GetByFilterOutput struct {
	UserID string
}
