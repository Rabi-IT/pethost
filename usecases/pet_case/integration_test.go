package pet_case_test

import (
	"fmt"
	"net/http"
	"pethost/fixtures"
	"pethost/frameworks/database/gateways/pet_gateway"
	"pethost/usecases/pet_case"
	"pethost/usecases/pet_case/pet"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func Test_Integration_should_create(t *testing.T) {
	fixtures.CleanDatabase()

	True := true
	Body := pet_case.CreateInput{
		Name:       "Pet",
		Gender:     pet.Male,
		Birthdate:  time.Now(),
		Weight:     pet.Medium,
		Species:    pet.Dog,
		Neutered:   &True,
		Vaccinated: &True,
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
	Body := map[string]any{
		"name":       "Pet Name",
		"tutorId":    user2,
		"birthDate":  time.Now(),
		"weight":     pet.Medium,
		"species":    pet.Dog,
		"gender":     pet.Male,
		"neutered":   true,
		"vaccinated": true,
	}

	petIdCreated := ""
	tokenUser1 := fixtures.User.Login(t, &user1)
	statusCode := fixtures.Post(t, fixtures.PostInput{
		Body:     Body,
		URI:      "/pet",
		Response: &petIdCreated,
		Token:    tokenUser1,
	})

	require.Equal(t, http.StatusCreated, statusCode, fmt.Sprintf("response: %s", petIdCreated))
	require.NotEmpty(t, petIdCreated)

	petsUser1, statusCode := fixtures.Pet.List(t, tokenUser1)
	require.Equal(t, http.StatusOK, statusCode)
	require.Len(t, petsUser1, 1)
	require.Equal(t, petIdCreated, petsUser1[0].ID, "user1 should have the pet created")

	tokenUser2 := fixtures.User.Login(t, &user2)
	petsUser2, statusCode := fixtures.Pet.List(t, tokenUser2)
	require.Equal(t, http.StatusOK, statusCode)
	require.Empty(t, petsUser2, "user2 should not have the pet created")
}

func Test_Integration_should_be_able_to_retrive_by_id(t *testing.T) {
	fixtures.CleanDatabase()

	token := fixtures.User.Login(t, nil)
	id := fixtures.Pet.Create(t, nil, token)

	responseBody := pet_gateway.GetByFilterOutput{}

	statusCode := fixtures.Get(t, fixtures.GetInput{
		URI:      "/pet/" + id,
		Response: &responseBody,
		Token:    token,
	})

	require.Equal(t, http.StatusOK, statusCode)

	EXPECTED := pet_gateway.GetByFilterOutput{
		Name:       "Name",
		Breed:      "Breed",
		Birthdate:  time.Date(2000, 0, 1, 0, 0, 0, 0, time.UTC),
		Gender:     pet.Male,
		Weight:     pet.Medium,
		Species:    "Species",
		Neutered:   true,
		Vaccinated: true,
	}

	require.Equal(t, EXPECTED, responseBody)
}

func Test_Integration_should_be_able_to_list(t *testing.T) {
	fixtures.CleanDatabase()

	token := fixtures.User.Login(t, nil)
	petId := fixtures.Pet.Create(t, nil, token)

	responseBody := []pet_gateway.ListOutput{}

	fixtures.Get(t, fixtures.GetInput{
		URI:      "/pet",
		Response: &responseBody,
		Token:    token,
	})

	EXPECTED := []pet_gateway.ListOutput{
		{
			ID:         petId,
			Name:       "Name",
			Breed:      "Breed",
			Birthdate:  time.Date(2000, 0, 1, 0, 0, 0, 0, time.UTC),
			Gender:     pet.Male,
			Weight:     pet.Medium,
			Species:    "Species",
			Neutered:   true,
			Vaccinated: true,
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

	statusCode := fixtures.Patch(t, fixtures.PatchInput{
		Body:  Body,
		URI:   "/pet/" + id,
		Token: token,
	})
	require.Equal(t, http.StatusOK, statusCode)

	found, statusCode := fixtures.Pet.GetByID(t, id, token)
	require.Equal(t, http.StatusOK, statusCode)

	EXPECTED := pet_gateway.GetByFilterOutput{
		Name:       "NewName",
		Breed:      "Breed",
		Birthdate:  time.Date(2000, 0, 1, 0, 0, 0, 0, time.UTC),
		Gender:     pet.Male,
		Weight:     pet.Medium,
		Species:    "Species",
		Neutered:   true,
		Vaccinated: true,
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

	EXPECTED := pet_gateway.GetByFilterOutput{}

	require.Equal(t, EXPECTED, found)
}
