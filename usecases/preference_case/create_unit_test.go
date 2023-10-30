package preference_case_test

import (
	"pethost/fixtures"
	"pethost/fixtures/mocks"
	"pethost/frameworks/database/gateways/preference_gateway"
	"pethost/usecases/preference_case"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func makeSut(gateway preference_gateway.PreferenceGateway) *preference_case.PreferenceCase {
	return preference_case.New(gateway)
}

func Test_Unit_Create__should_not_fail_if_required_fields_are_filled_in(t *testing.T) {
	gateway := mocks.NewPreferenceGateway(t)
	expectedID := "ANY_ID"
	gateway.On("Create", mock.Anything).Return(expectedID, nil)
	sut := makeSut(gateway)

	False := false
	id, err := sut.Create(fixtures.DUMMY_CONTEXT, &preference_case.CreateInput{
		OnlyVaccinated:          &False,
		AcceptElderly:           &False,
		AcceptOnlyNeuteredMales: &False,
		AcceptFemales:           &False,
		DaysOfMonth:             (1 << 32) - 1,
		AcceptFemaleInHeat:      &False,
		AcceptPuppies:           &False,
		AcceptMales:             &False,
		PetWeight:               (1 << 5) - 1,
	})

	require.Nil(t, err)
	require.Equal(t, expectedID, id)
}
