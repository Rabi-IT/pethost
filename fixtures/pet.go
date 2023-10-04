package fixtures

import (
	"net/http"
	"pethost/adapters/gateways/pet_gateway"
	"pethost/usecases/pet_case"
	"testing"

	"github.com/stretchr/testify/require"
)

type petFixture struct{}

var Pet = petFixture{}

func (petFixture) Create(t *testing.T, input *pet_case.CreateInput) string {
	Body := input
	if Body == nil {
		Body = &pet_case.CreateInput{
			Name:      "Name",
			Breed:     "Breed",
			Size:      "Size",
			Birthdate: "Birthdate",
			Gender:    "Gender",
			Weight:    "Weight",
			Species:   "Species",
		}
	}

	id := ""
	statusCode := Post(t, PostInput{
		Body:     Body,
		URI:      "/pet",
		Response: &id,
	})

	require.Equal(t, http.StatusCreated, statusCode)
	require.NotEmpty(t, id)

	return id
}

func (petFixture) GetByID(t *testing.T, id string) (pet_gateway.GetByIDOutput, int) {
	found := pet_gateway.GetByIDOutput{}

	input := GetInput{
		URI:      "/pet/" + id,
		Response: &found,
	}

	statusCode := Get(t, input)

	return found, statusCode
}
