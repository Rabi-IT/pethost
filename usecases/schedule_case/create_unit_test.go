package schedule_case_test

import (
	"pethost/fixtures"
	"pethost/fixtures/mocks"
	"pethost/frameworks/database/gateways/pet_gateway"
	"pethost/usecases/pet_case"
	"pethost/usecases/pet_case/pet"
	"pethost/usecases/preference_case"
	"pethost/usecases/schedule_case"
	"testing"
	"time"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

type makeSutOutput struct {
	sut               *schedule_case.ScheduleCase
	gateway           *mocks.ScheduleGateway
	petGateway        *mocks.PetGateway
	preferenceGateway *mocks.PreferenceGateway
}

func makeSut(t *testing.T) *makeSutOutput {
	gateway := mocks.NewScheduleGateway(t)
	preferenceGateway := mocks.NewPreferenceGateway(t)
	petGateway := mocks.NewPetGateway(t)

	sut := schedule_case.New(
		gateway,
		preference_case.New(preferenceGateway),
		pet_case.New(petGateway),
	)

	return &makeSutOutput{
		sut:               sut,
		gateway:           gateway,
		petGateway:        petGateway,
		preferenceGateway: preferenceGateway,
	}
}

func Test_Unit(t *testing.T) {
	t.Skip()
	test := makeSut(t)
	test.petGateway.On("GetByID", mock.Anything, &pet_gateway.GetByFilterOutput{
		Name:      "A beautyful name",
		Breed:     "A great breed",
		Birthdate: time.Date(2000, 0, 1, 0, 0, 0, 0, time.UTC),
		Gender:    pet.Male,
		Weight:    fixtures.Preference.AllPetWeight,
		Species:   string(pet.Dog),
		Neutered:  true,
	})

	id, err := test.sut.Create(fixtures.DUMMY_CONTEXT, &schedule_case.CreateInput{
		PetIDs: []string{"ANY_ID"},
		HostID: "ANY_ID",
		Notes:  "ANY_ID",
	})

	require.Nil(t, err)

	EXPECTED := "EXPECTED_ID"
	require.Equal(t, EXPECTED, id)
}
