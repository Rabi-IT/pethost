package preference_case_test

import (
	"net/http"
	"pethost/fixtures"
	"strings"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func Test_Integration_Create__should_be_able_to_create(t *testing.T) {
	fixtures.CleanDatabase()

	Body := &map[string]any{
		"onlyVaccinated":          false,
		"acceptElderly":           false,
		"acceptOnlyNeuteredMales": false,
		"acceptFemales":           false,
		"daysOfMonth":             fixtures.Preference.AllDaysOfMonth,
		"acceptFemaleInHeat":      false,
		"acceptPuppies":           false,
		"acceptMales":             false,
		"petWeight":               fixtures.Preference.AllPetWeight,
	}

	id := ""
	token := fixtures.User.Login(t, nil)
	statusCode := fixtures.Post(t, fixtures.PostInput{
		Body:     Body,
		URI:      fixtures.Preference.URI,
		Response: &id,
		Token:    token,
	})

	if _, err := uuid.Parse(id); err != nil {
		require.Fail(t, "response error", id)
	}
	require.Equal(t, http.StatusCreated, statusCode)
}

func Test_Integration_Create__should_fail_if_required_fields_are_empty(t *testing.T) {
	fixtures.CleanDatabase()

	Body := &map[string]any{
		"onlyVaccinated":          nil,
		"acceptElderly":           nil,
		"acceptOnlyNeuteredMales": nil,
		"acceptFemales":           nil,
		"daysOfMonth":             nil,
		"acceptFemaleInHeat":      nil,
		"acceptPuppies":           nil,
		"acceptMales":             nil,
		"petWeight":               nil,
	}

	errorResponse := ""
	token := fixtures.User.Login(t, nil)
	statusCode := fixtures.Post(t, fixtures.PostInput{
		Body:     Body,
		URI:      fixtures.Preference.URI,
		Response: &errorResponse,
		Token:    token,
	})

	expectedMsg := []string{
		"Key: 'CreateInput.OnlyVaccinated' Error:Field validation for 'OnlyVaccinated' failed on the 'required' tag",
		"Key: 'CreateInput.AcceptElderly' Error:Field validation for 'AcceptElderly' failed on the 'required' tag",
		"Key: 'CreateInput.AcceptOnlyNeuteredMales' Error:Field validation for 'AcceptOnlyNeuteredMales' failed on the 'required' tag",
		"Key: 'CreateInput.AcceptFemales' Error:Field validation for 'AcceptFemales' failed on the 'required' tag",
		"Key: 'CreateInput.DaysOfMonth' Error:Field validation for 'DaysOfMonth' failed on the 'required' tag",
		"Key: 'CreateInput.AcceptFemaleInHeat' Error:Field validation for 'AcceptFemaleInHeat' failed on the 'required' tag",
		"Key: 'CreateInput.AcceptPuppies' Error:Field validation for 'AcceptPuppies' failed on the 'required' tag",
		"Key: 'CreateInput.AcceptMales' Error:Field validation for 'AcceptMales' failed on the 'required' tag",
		"Key: 'CreateInput.PetWeight' Error:Field validation for 'PetWeight' failed on the 'required' tag",
	}

	require.Equal(t, strings.Join(expectedMsg, "\n"), errorResponse)
	require.Equal(t, http.StatusBadRequest, statusCode)
}
