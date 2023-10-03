package fixtures

import (
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

	response := RawPost(t, RawPostInput{
		Body: Body,
		URI:  "/pet",
	})

	require.NotEmpty(t, response)

	return response
}

func (petFixture) GetByID(t *testing.T, id string) pet_gateway.GetByIDOutput {
	found := pet_gateway.GetByIDOutput{}

	Get(t, GetInput{
		URI:      "/pet/" + id,
		Response: &found,
	})

	return found
}
