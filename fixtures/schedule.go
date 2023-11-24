package fixtures

import (
	"fmt"
	"net/http"
	"pethost/usecases/schedule_case"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

type scheduleFixture struct {
	URI string
}

var Schedule = scheduleFixture{
	URI: "/schedule/",
}

func (scheduleFixture) Create(t *testing.T, input schedule_case.CreateInput, hostToken string) (string, int) {
	id := ""

	statusCode := Post(t, PostInput{
		Body:     input,
		URI:      Schedule.URI,
		Response: &id,
		Token:    hostToken,
	})

	return id, statusCode
}

type ScheduleDefaultOutput struct {
	PreferenceDefaultOutput
	ID        string
	Status    int
	StartDate time.Time
	EndDate   time.Time
}

func (s scheduleFixture) CreateDefault(t *testing.T) (output ScheduleDefaultOutput) {
	scenario := Preference.CreateDefault(t, nil)

	startDate := time.Now().AddDate(0, 3, 0)
	endDate := time.Now().AddDate(0, 10, 0)

	scheduleId, statusCode := s.Create(t, schedule_case.CreateInput{
		HostID:        scenario.HostID,
		PetIDs:        []string{scenario.PetID},
		StartDate:     startDate,
		EndDate:       endDate,
		FemalesInHeat: nil,
		Notes:         "",
	}, scenario.TutorToken)

	require.Equal(t, http.StatusCreated, statusCode, fmt.Sprintf("response: %s", scheduleId))

	output.ID = scheduleId
	output.Status = statusCode
	output.StartDate = startDate
	output.EndDate = endDate

	output.PreferenceDefaultOutput = scenario

	return
}
