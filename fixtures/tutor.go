package fixtures

import (
	"net/http"
	"pethost/adapters/gateways/tutor_gateway"
	"pethost/config"
	"pethost/usecases/tutor_case"
	"testing"

	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/require"
)

type tutorFixture struct {
	URI string
}

var Tutor = tutorFixture{"/tutor/"}

func (tutorFixture) Create(t *testing.T, input *tutor_case.CreateInput) string {
	Body := input
	if Body == nil {
		Body = &tutor_case.CreateInput{
			Name:           "Name",
			Photo:          "Photo",
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
		URI:      Tutor.URI,
		Response: &id,
	})

	require.Equal(t, http.StatusCreated, statusCode)
	require.NotEmpty(t, id)

	return id
}

func (tutorFixture) GetByID(t *testing.T, id string) (tutor_gateway.GetByIDOutput, int) {
	found := tutor_gateway.GetByIDOutput{}

	input := GetInput{
		URI:      Tutor.URI + id,
		Response: &found,
	}

	statusCode := Get(t, input)

	return found, statusCode
}

func (tutorFixture) Login(t *testing.T, id *string) string {
	tutorId := ""
	tutor := tutor_case.CreateInput{}
	if id == nil {
		tutor = tutor_case.CreateInput{
			Name:           "Tutor Test",
			Photo:          "Photo",
			TaxID:          "TaxID",
			City:           "City",
			State:          "State",
			Phone:          "Phone",
			ZIP:            "ZIP",
			SocialID:       "SocialID",
			Email:          "tutor_test@email.com",
			EmergencyPhone: "EmergencyPhone",
			Neighborhood:   "Neighborhood",
			Street:         "Street",
			Complement:     "Complement",
		}

		tutorId = Tutor.Create(t, &tutor)
	} else {
		retrievedTutor, statusCode := Tutor.GetByID(t, *id)
		require.Equal(t, http.StatusOK, statusCode)
		tutor.Name = retrievedTutor.Name
		tutor.Email = retrievedTutor.Email
		tutorId = *id
	}

	claims := jwt.MapClaims{
		"user_id": tutorId,
		"name":    tutor.Name,
		"email":   tutor.Email,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tk, err := token.SignedString([]byte(config.AuthSecret))
	require.Nil(t, err)

	return tk
}
