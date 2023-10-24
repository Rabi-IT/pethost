package pet_case_test

import (
	"net/http"
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

	id := ""
	token := fixtures.Tutor.Login(t, nil)
	statusCode := fixtures.Post(t, fixtures.PostInput{
		Body:     Body,
		URI:      "/pet",
		Response: &id,
		Token:    token,
	})

	require.Equal(t, http.StatusCreated, statusCode)
	require.NotEmpty(t, id)
}

func Test_Integration_should_be_able_to_retrive_by_id(t *testing.T) {
	fixtures.CleanDatabase()

	token := fixtures.Tutor.Login(t, nil)
	id := fixtures.Pet.Create(t, nil, token)

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

	token := fixtures.Tutor.Login(t, nil)
	fixtures.Pet.Create(t, nil, token)

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

	token := fixtures.Tutor.Login(t, nil)
	id := fixtures.Pet.Create(t, nil, token)

	Body := pet_case.PatchValues{
		Name: "NewName",
	}

	ok := fixtures.Patch(t, fixtures.PatchInput{
		Body: Body,
		URI:  "/pet/" + id,
	})
	require.True(t, ok == "OK")

	found, statusCode := fixtures.Pet.GetByID(t, id)
	require.Equal(t, http.StatusOK, statusCode)

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

func Test_Integration_should_be_able_to_delete(t *testing.T) {
	fixtures.CleanDatabase()

	token := fixtures.Tutor.Login(t, nil)
	id := fixtures.Pet.Create(t, nil, token)

	Body := pet_case.PatchValues{
		Name: "NewName",
	}

	respBody, statusCode := fixtures.Delete(t, fixtures.DeleteInput{
		Body: Body,
		URI:  "/pet/" + id,
	})

	require.Equal(t, statusCode, http.StatusNoContent)
	require.Equal(t, respBody, "")

	found, statusCode := fixtures.Pet.GetByID(t, id)
	require.Equal(t, statusCode, http.StatusNotFound)

	EXPECTED := pet_gateway.GetByIDOutput{}

	require.Equal(t, EXPECTED, found)
}
