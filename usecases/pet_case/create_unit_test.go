package pet_case_test

import (
	"pethost/fixtures"
	"pethost/fixtures/mocks"
	"pethost/frameworks/database/gateways/pet_gateway"
	"pethost/usecases/pet_case"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func makeSut(g pet_gateway.PetGateway) *pet_case.PetCase {
	return pet_case.New(g)
}

func Test_Unit_Create__should_fail_if_name_is_empty(t *testing.T) {
	sut := makeSut(nil)

	_, err := sut.Create(fixtures.DUMMY_CONTEXT, &pet_case.CreateInput{
		Name:      "",
		Breed:     "Breed",
		Size:      "Size",
		Birthdate: "Birthdate",
		Gender:    "Gender",
		Weight:    "Weight",
		Species:   "Species",
	})

	expectedMsg := "Key: 'CreateInput.Name' Error:Field validation for 'Name' failed on the 'required' tag"
	require.Equal(t, expectedMsg, err.Error())
}

func Test_Unit_Create__should_not_fail_if_all_optional_fields_are_not_filled_in(t *testing.T) {
	gateway := mocks.NewPetGateway(t)
	expectedID := "ANY_ID"
	gateway.On("Create", mock.Anything).Return(expectedID, nil)
	sut := pet_case.New(gateway)

	id, err := sut.Create(fixtures.DUMMY_CONTEXT, &pet_case.CreateInput{
		Name:      "Name",
		Breed:     "",
		Size:      "",
		Birthdate: "",
		Gender:    "",
		Weight:    "",
		Species:   "",
	})

	require.Nil(t, err)
	require.Equal(t, expectedID, id)
}
