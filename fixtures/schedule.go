package fixtures

import (
	"pethost/usecases/schedule_case"
	"testing"
)

type scheduleFixture struct {
	URI string
}

var Schedule = scheduleFixture{
	URI: "/schedule/",
}

func (scheduleFixture) Create(t *testing.T, input schedule_case.CreateInput, hostToken string) string {
	id := ""

	Post(t, PostInput{
		Body:     input,
		URI:      Schedule.URI,
		Response: &id,
		Token:    hostToken,
	})

	return id
}
