package fixtures

import (
	"net/http"
	"pethost/frameworks/database/gateways/pet_gateway"
	"pethost/usecases/pet_case"
	"pethost/usecases/pet_case/pet"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

type petFixture struct {
	URI string
}

var Pet = petFixture{
	URI: "/pet/",
}

func (petFixture) Create(t *testing.T, input *pet_case.CreateInput, token string) string {
	Body := input
	if Body == nil {
		True := true
		Body = &pet_case.CreateInput{
			Name:       "Name",
			Breed:      "Breed",
			Birthdate:  time.Date(2000, 0, 1, 0, 0, 0, 0, time.UTC),
			Gender:     pet.Male,
			Weight:     pet.Medium,
			Species:    "Species",
			Neutered:   &True,
			Vaccinated: &True,
		}
	}

	id := ""
	statusCode := Post(t, PostInput{
		Body:     Body,
		URI:      Pet.URI,
		Response: &id,
		Token:    token,
	})

	require.Equal(t, http.StatusCreated, statusCode)
	require.NotEmpty(t, id)

	return id
}

func (petFixture) GetByID(t *testing.T, id string, token string) (pet_gateway.GetByFilterOutput, int) {
	found := pet_gateway.GetByFilterOutput{}

	input := GetInput{
		URI:      Pet.URI + id,
		Response: &found,
		Token:    token,
	}

	statusCode := Get(t, input)

	return found, statusCode
}

func (petFixture) List(t *testing.T, token string) ([]pet_gateway.ListOutput, int) {
	found := []pet_gateway.ListOutput{}

	input := GetInput{
		URI:      Pet.URI,
		Response: &found,
		Token:    token,
	}

	statusCode := Get(t, input)

	return found, statusCode
}

func (petFixture) Patch(t *testing.T, petId string, newValues pet_case.PatchValues, tutorToken string) (string, int) {
	response := ""
	statusCode := Patch(t, PatchInput{
		Response: &response,
		URI:      "/pet/" + petId,
		Body:     newValues,
		Token:    tutorToken,
	})

	return response, statusCode
}
