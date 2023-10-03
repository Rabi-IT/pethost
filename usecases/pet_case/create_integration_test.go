package pet_case_test

import (
	"pethost/adapters/gateways/pet_gateway"
	"pethost/fixtures"
	"pethost/usecases/pet_case"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Integration_should_create(t *testing.T) {
	fixtures.CleanDatabase()

	Body := pet_case.CreateInput{
		Name: "Pet",
	}

	response := fixtures.RawPost(t, fixtures.RawPostInput{
		Body: Body,
		URI:  "/pet",
	})

	require.NotEmpty(t, response)
}

func Test_Integration_should_be_able_to_retrive_by_id(t *testing.T) {
	fixtures.CleanDatabase()

	id := fixtures.Pet.Create(t, nil)

	responseBody := pet_gateway.GetByIDOutput{}

	fixtures.Get(t, fixtures.GetInput{
		URI:      "/pet/" + id,
		Response: &responseBody,
	})

	EXPECTED := pet_gateway.GetByIDOutput{
		Name:      "Name",
		Breed:     "Breed",
		Size:      "Size",
		Birthdate: "Birthdate",
		Gender:    "Gender",
		Weight:    "Weight",
		Species:   "Species",
	}

	require.Equal(t, EXPECTED, responseBody)
}

func Test_Integration_should_be_able_to_list(t *testing.T) {
	fixtures.CleanDatabase()

	fixtures.Pet.Create(t, nil)

	responseBody := []pet_gateway.ListOutput{}

	fixtures.Get(t, fixtures.GetInput{
		URI:      "/pet",
		Response: &responseBody,
	})

	EXPECTED := []pet_gateway.ListOutput{
		{
			Name:      "Name",
			Breed:     "Breed",
			Size:      "Size",
			Birthdate: "Birthdate",
			Gender:    "Gender",
			Weight:    "Weight",
			Species:   "Species",
		},
	}

	require.Equal(t, EXPECTED, responseBody)
}

func Test_Integration_should_be_able_to_update(t *testing.T) {
	fixtures.CleanDatabase()

	id := fixtures.Pet.Create(t, nil)

	Body := pet_case.PatchValues{
		Name: "NewName",
	}

	ok := fixtures.Patch(t, fixtures.PatchInput{
		Body: Body,
		URI:  "/pet/" + id,
	})
	require.True(t, ok == "OK")

	found := fixtures.Pet.GetByID(t, id)

	EXPECTED := pet_gateway.GetByIDOutput{
		Name:      "NewName",
		Breed:     "Breed",
		Size:      "Size",
		Birthdate: "Birthdate",
		Gender:    "Gender",
		Weight:    "Weight",
		Species:   "Species",
	}

	require.Equal(t, EXPECTED, found)
}
