package pet_case_test

import (
	"net/http"
	"pethost/fixtures"
	"pethost/frameworks/database/gateways/pet_gateway"
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
	token := fixtures.User.Login(t, nil)
	statusCode := fixtures.Post(t, fixtures.PostInput{
		Body:     Body,
		URI:      "/pet",
		Response: &id,
		Token:    token,
	})

	require.Equal(t, http.StatusCreated, statusCode)
	require.NotEmpty(t, id)
}

func Test_Integration_should_not_be_able_to_create_pet_to_another_user(t *testing.T) {
	fixtures.CleanDatabase()

	user1 := fixtures.User.Create(t, nil)
	user2 := fixtures.User.Create(t, nil)
	Body := map[string]string{
		"name":    "Pet Name",
		"tutorId": user2,
	}

	id := ""
	tokenUser1 := fixtures.User.Login(t, &user1)
	statusCode := fixtures.Post(t, fixtures.PostInput{
		Body:     Body,
		URI:      "/pet",
		Response: &id,
		Token:    tokenUser1,
	})

	require.Equal(t, http.StatusCreated, statusCode)
	require.NotEmpty(t, id)

	petsUser1, statusCode := fixtures.Pet.List(t, tokenUser1)
	require.Equal(t, http.StatusOK, statusCode)
	expected := []pet_gateway.ListOutput{{Name: "Pet Name"}}
	require.Equal(t, expected, petsUser1)

	tokenUser2 := fixtures.User.Login(t, &user2)
	petsUser2, statusCode := fixtures.Pet.List(t, tokenUser2)
	require.Equal(t, http.StatusOK, statusCode)
	require.Empty(t, petsUser2)
}

func Test_Integration_should_be_able_to_retrive_by_id(t *testing.T) {
	fixtures.CleanDatabase()

	token := fixtures.User.Login(t, nil)
	id := fixtures.Pet.Create(t, nil, token)

	responseBody := pet_gateway.GetByIDOutput{}

	statusCode := fixtures.Get(t, fixtures.GetInput{
		URI:      "/pet/" + id,
		Response: &responseBody,
		Token:    token,
	})

	require.Equal(t, http.StatusOK, statusCode)

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

	token := fixtures.User.Login(t, nil)
	fixtures.Pet.Create(t, nil, token)

	responseBody := []pet_gateway.ListOutput{}

	fixtures.Get(t, fixtures.GetInput{
		URI:      "/pet",
		Response: &responseBody,
		Token:    token,
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

	token := fixtures.User.Login(t, nil)
	id := fixtures.Pet.Create(t, nil, token)

	Body := pet_case.PatchValues{
		Name: "NewName",
	}

	ok := fixtures.Patch(t, fixtures.PatchInput{
		Body:  Body,
		URI:   "/pet/" + id,
		Token: token,
	})
	require.True(t, ok == "OK")

	found, statusCode := fixtures.Pet.GetByID(t, id, token)
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

	token := fixtures.User.Login(t, nil)
	id := fixtures.Pet.Create(t, nil, token)

	Body := pet_case.PatchValues{
		Name: "NewName",
	}

	respBody, statusCode := fixtures.Delete(t, fixtures.DeleteInput{
		Body:  Body,
		URI:   "/pet/" + id,
		Token: token,
	})

	require.Equal(t, statusCode, http.StatusNoContent)
	require.Equal(t, respBody, "")

	found, statusCode := fixtures.Pet.GetByID(t, id, token)
	require.Equal(t, statusCode, http.StatusNotFound)

	EXPECTED := pet_gateway.GetByIDOutput{}

	require.Equal(t, EXPECTED, found)
}
