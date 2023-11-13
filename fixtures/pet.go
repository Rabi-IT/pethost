package fixtures

import (
	"net/http"
	"pethost/frameworks/database/gateways/pet_gateway"
	"pethost/usecases/pet_case"
	"testing"

	"github.com/stretchr/testify/require"
)

type petFixture struct {
	URI          string
	VerySmallPet uint8
	SmallPet     uint8
	MediumPet    uint8
	LargePet     uint8
	VeryLargePet uint8
}

var Pet = petFixture{
	URI:          "/pet/",
	VerySmallPet: 0b00001,
	SmallPet:     0b00010,
	MediumPet:    0b00100,
	LargePet:     0b01000,
	VeryLargePet: 0b10000,
}

func (petFixture) Create(t *testing.T, input *pet_case.CreateInput, token string) string {
	Body := input
	if Body == nil {
		Body = &pet_case.CreateInput{
			Name:      "Name",
			Breed:     "Breed",
			Birthdate: "Birthdate",
			Gender:    "Gender",
			Weight:    Pet.MediumPet,
			Species:   "Species",
			Neutered:  true,
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

func (petFixture) GetByID(t *testing.T, id string, token string) (pet_gateway.GetByIDOutput, int) {
	found := pet_gateway.GetByIDOutput{}

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
