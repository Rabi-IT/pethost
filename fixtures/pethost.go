package fixtures

import (
	"net/http"
	"pethost/adapters/gateways/pethost_gateway"
	"pethost/usecases/pethost_case"
	"testing"

	"github.com/stretchr/testify/require"
)

const baseURI = "/pethost/"

type pethostFixture struct{}

var Pethost = pethostFixture{}

func (pethostFixture) Create(t *testing.T, input *pethost_case.CreateInput) string {
	Body := input
	if Body == nil {
		Body = &pethost_case.CreateInput{
			Name:           "Name",
			TaxID:          "TaxID",
			City:           "City",
			State:          "State",
			Phone:          "Phone",
			ZIP:            "ZIP",
			SocialID:       "SocialID",
			Email:          "Email",
			EmergencyPhone: "EmergencyPhone",
			Neighborhood:   "Neighborhood",
			Street:         "Street",
			Complement:     "Complement",
		}
	}

	id := ""
	statusCode := Post(t, PostInput{
		Body:     Body,
		URI:      baseURI,
		Response: &id,
	})

	require.Equal(t, http.StatusCreated, statusCode)
	require.NotEmpty(t, id)

	return id
}

func (pethostFixture) GetByID(t *testing.T, id string) (pethost_gateway.GetByIDOutput, int) {
	found := pethost_gateway.GetByIDOutput{}

	input := GetInput{
		URI:      baseURI + id,
		Response: &found,
	}

	statusCode := Get(t, input)

	return found, statusCode
}
