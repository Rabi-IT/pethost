package pet_case_test

import (
	"pethost/fixtures"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Create__should_save(t *testing.T) {
	// fixtures.CleanDatabase()

	Body := map[string]string{
		"Name": "Pet",
	}

	response := fixtures.Post(t, fixtures.PostInput{
		Body: Body,
		URI:  "/pet",
	})

	require.NotEmpty(t, response)
}
