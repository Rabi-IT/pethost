package fixtures

import (
	"net/http"
	"pethost/usecases/preference_case"
	"testing"

	"github.com/stretchr/testify/require"
)

type preferenceFixture struct {
	URI string
}

var Preference = preferenceFixture{"/preference/"}

func (preferenceFixture) Create(t *testing.T, input *preference_case.CreateInput) string {
	Body := input
	if Body == nil {
		False := false
		Body = &preference_case.CreateInput{
			OnlyVaccinated:          &False,
			AcceptElderly:           &False,
			AcceptOnlyNeuteredMales: &False,
			AcceptFemales:           &False,
			DaysOfMonth:             (1 << 32) - 1,
			AcceptFemaleInHeat:      &False,
			AcceptPuppies:           &False,
			AcceptMales:             &False,
			PetWeight:               (1 << 5) - 1,
		}
	}

	id := ""
	statusCode := Post(t, PostInput{
		Body:     Body,
		URI:      Preference.URI,
		Response: &id,
		Token:    SystemToken(t),
	})

	require.Equal(t, http.StatusCreated, statusCode)
	require.NotEmpty(t, id)

	return id
}
