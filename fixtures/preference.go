package fixtures

import (
	"pethost/usecases/preference_case"
	"testing"
)

type preferenceFixture struct {
	URI            string
	AllDaysOfMonth uint32
	AllPetWeight   uint8
	OnlyLargePets  uint8
	OnlySmallPets  uint8
}

var Preference = preferenceFixture{
	URI:            "/preference/",
	AllDaysOfMonth: (1 << 32) - 1,
	AllPetWeight:   (1 << 5) - 1,
	OnlyLargePets:  0b11000,
	OnlySmallPets:  0b00011,
}

func (preferenceFixture) Create(t *testing.T, token string, input *preference_case.CreateInput) string {
	Body := input
	if Body == nil {
		False := false
		Body = &preference_case.CreateInput{
			OnlyVaccinated:          &False,
			AcceptElderly:           &False,
			AcceptOnlyNeuteredMales: &False,
			AcceptFemales:           &False,
			DaysOfMonth:             Preference.AllDaysOfMonth,
			AcceptFemaleInHeat:      &False,
			AcceptPuppies:           &False,
			AcceptMales:             &False,
			PetWeight:               Preference.AllPetWeight,
		}
	}

	id := ""
	Post(t, PostInput{
		Body:     Body,
		URI:      Preference.URI,
		Response: &id,
		Token:    token,
	})

	return id
}

type CreateDefaultOutput struct {
	TutorID      string
	TutorToken   string
	HostID       string
	HostToken    string
	PetID        string
	PreferenceID string
}

func (preferenceFixture) CreateDefault(t *testing.T, input *preference_case.CreateInput) (output CreateDefaultOutput) {
	output.TutorID = User.Create(t, nil)
	output.TutorToken = User.Login(t, &output.TutorID)
	output.PetID = Pet.Create(t, nil, output.TutorToken)

	output.HostID = User.Create(t, nil)
	output.HostToken = User.Login(t, &output.HostID)

	if input == nil {
		True := true
		input = &preference_case.CreateInput{
			OnlyVaccinated:          &True,
			AcceptElderly:           &True,
			AcceptOnlyNeuteredMales: &True,
			AcceptFemales:           &True,
			DaysOfMonth:             Preference.AllDaysOfMonth,
			AcceptFemaleInHeat:      &True,
			AcceptPuppies:           &True,
			AcceptMales:             &True,
			PetWeight:               Preference.AllPetWeight,
		}
	}

	output.PreferenceID = Preference.Create(t, output.HostToken, input)

	return
}
