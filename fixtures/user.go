package fixtures

import (
	"net/http"
	"pethost/adapters/gateways/user_gateway"
	"pethost/config"
	"pethost/usecases/user_case"
	"testing"

	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/require"
)

type userFixture struct {
	URI string
}

var User = userFixture{"/user/"}

func (userFixture) Create(t *testing.T, input *user_case.CreateInput) string {
	Body := input
	if Body == nil {
		Body = &user_case.CreateInput{
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
		URI:      User.URI,
		Response: &id,
	})

	require.Equal(t, http.StatusCreated, statusCode)
	require.NotEmpty(t, id)

	return id
}

func (userFixture) GetByID(t *testing.T, id string) (user_gateway.GetByIDOutput, int) {
	found := user_gateway.GetByIDOutput{}

	input := GetInput{
		URI:      User.URI + id,
		Response: &found,
	}

	statusCode := Get(t, input)

	return found, statusCode
}

func (userFixture) Login(t *testing.T, id *string) string {
	userId := ""
	user := user_case.CreateInput{}
	if id == nil {
		user = user_case.CreateInput{
			Name:           "User Test",
			Photo:          "Photo",
			TaxID:          "TaxID",
			City:           "City",
			State:          "State",
			Phone:          "Phone",
			ZIP:            "ZIP",
			SocialID:       "SocialID",
			Email:          "user_test@email.com",
			EmergencyPhone: "EmergencyPhone",
			Neighborhood:   "Neighborhood",
			Street:         "Street",
			Complement:     "Complement",
		}

		userId = User.Create(t, &user)
	} else {
		retrievedUser, statusCode := User.GetByID(t, *id)
		require.Equal(t, http.StatusOK, statusCode)
		user.Name = retrievedUser.Name
		user.Email = retrievedUser.Email
		userId = *id
	}

	claims := jwt.MapClaims{
		"user_id": userId,
		"name":    user.Name,
		"email":   user.Email,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tk, err := token.SignedString([]byte(config.AuthSecret))
	require.Nil(t, err)

	return tk
}
