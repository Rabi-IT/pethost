package schedule_case_test

import (
	"fmt"
	"net/http"
	"pethost/fixtures"
	"pethost/frameworks/database/gateways/schedule_gateway"
	"pethost/usecases/auth_case/role"
	"pethost/usecases/schedule_case"
	"pethost/usecases/schedule_case/schedule_status"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func Test_Integration_Patch(t *testing.T) {
	type testCase struct {
		title          string
		expectedStatus int
		seed           func() (response string, statusCode int, scenario fixtures.ScheduleDefaultOutput)
		expected       func(scenario fixtures.ScheduleDefaultOutput) schedule_gateway.PaginateData
	}

	tests := []testCase{
		{
			title:          fmt.Sprintf("should be able to patch status to %s if status is %s", schedule_status.Accepted, schedule_status.Open),
			expectedStatus: http.StatusOK,
			seed: func() (response string, statusCode int, scenario fixtures.ScheduleDefaultOutput) {
				scenario = fixtures.Schedule.CreateDefault(t)

				statusCode = fixtures.Patch(t, fixtures.PatchInput{
					URI: fixtures.Schedule.URI + scenario.ID,
					Body: schedule_case.PatchValues{
						Status: schedule_status.Accepted,
					},
					Response: &response,
					Token:    scenario.HostToken,
				})

				return
			},
			expected: func(scenario fixtures.ScheduleDefaultOutput) schedule_gateway.PaginateData {
				return schedule_gateway.PaginateData{
					Notes:     "",
					StartDate: scenario.StartDate,
					EndDate:   scenario.EndDate,
					TutorID:   scenario.TutorID,
					PetIDs:    []string{scenario.PetID},
					Status:    schedule_status.Accepted,
				}
			},
		},

		{
			title:          fmt.Sprintf("should be able to patch status to %s if status is %s", schedule_status.Paid, schedule_status.Accepted),
			expectedStatus: http.StatusOK,
			seed: func() (response string, statusCode int, scenario fixtures.ScheduleDefaultOutput) {
				scenario = fixtures.Schedule.CreateDefault(t)

				statusCode = fixtures.Patch(t, fixtures.PatchInput{
					URI: fixtures.Schedule.URI + scenario.ID,
					Body: schedule_case.PatchValues{
						Status: schedule_status.Accepted,
					},
					Response: &response,
					Token:    scenario.HostToken,
				})

				require.Equal(t, http.StatusOK, statusCode)

				statusCode = fixtures.Patch(t, fixtures.PatchInput{
					URI: fixtures.Schedule.URI + scenario.ID,
					Body: schedule_case.PatchValues{
						Status: schedule_status.Paid,
					},
					Response: &response,
					Token:    fixtures.BackofficeToken(t),
				})

				return
			},
			expected: func(scenario fixtures.ScheduleDefaultOutput) schedule_gateway.PaginateData {
				return schedule_gateway.PaginateData{
					StartDate: scenario.StartDate,
					EndDate:   scenario.EndDate,
					Notes:     "",
					TutorID:   scenario.TutorID,
					PetIDs:    []string{scenario.PetID},
					Status:    schedule_status.PaidAccepted,
				}
			},
		},

		{
			title:          fmt.Sprintf("should be able to patch status to %s if status is %s", schedule_status.Canceled, schedule_status.Open),
			expectedStatus: http.StatusOK,
			seed: func() (response string, statusCode int, scenario fixtures.ScheduleDefaultOutput) {
				scenario = fixtures.Schedule.CreateDefault(t)

				statusCode = fixtures.Patch(t, fixtures.PatchInput{
					URI: fixtures.Schedule.URI + scenario.ID,
					Body: schedule_case.PatchValues{
						Status: schedule_status.Canceled,
					},
					Response: &response,
					Token:    scenario.HostToken,
				})

				return
			},
			expected: func(scenario fixtures.ScheduleDefaultOutput) schedule_gateway.PaginateData {
				return schedule_gateway.PaginateData{
					StartDate: scenario.StartDate,
					EndDate:   scenario.EndDate,
					Notes:     "",
					TutorID:   scenario.TutorID,
					PetIDs:    []string{scenario.PetID},
					Status:    schedule_status.Canceled,
				}
			},
		},

		{
			title:          fmt.Sprintf("should not be able to patch status to %s if status is %s", schedule_status.Accepted, schedule_status.Canceled),
			expectedStatus: http.StatusNotFound,
			seed: func() (response string, statusCode int, scenario fixtures.ScheduleDefaultOutput) {
				scenario = fixtures.Schedule.CreateDefault(t)

				statusCode = fixtures.Patch(t, fixtures.PatchInput{
					URI: fixtures.Schedule.URI + scenario.ID,
					Body: schedule_case.PatchValues{
						Status: schedule_status.Canceled,
					},
					Response: &response,
					Token:    scenario.HostToken,
				})

				require.Equal(t, http.StatusOK, statusCode)

				statusCode = fixtures.Patch(t, fixtures.PatchInput{
					URI: fixtures.Schedule.URI + scenario.ID,
					Body: schedule_case.PatchValues{
						Status: schedule_status.Accepted,
					},
					Response: &response,
					Token:    scenario.HostToken,
				})

				return
			},
			expected: func(scenario fixtures.ScheduleDefaultOutput) schedule_gateway.PaginateData {
				return schedule_gateway.PaginateData{
					StartDate: scenario.StartDate,
					EndDate:   scenario.EndDate,
					Notes:     "",
					TutorID:   scenario.TutorID,
					PetIDs:    []string{scenario.PetID},
					Status:    schedule_status.Canceled,
				}
			},
		},

		{
			title:          fmt.Sprintf("should be able to patch status to %s if user role is %s", schedule_status.Paid, role.Backoffice),
			expectedStatus: http.StatusOK,
			seed: func() (response string, statusCode int, scenario fixtures.ScheduleDefaultOutput) {
				scenario = fixtures.Schedule.CreateDefault(t)

				statusCode = fixtures.Patch(t, fixtures.PatchInput{
					URI: fixtures.Schedule.URI + scenario.ID,
					Body: schedule_case.PatchValues{
						Status: schedule_status.Paid,
					},
					Response: &response,
					Token:    fixtures.BackofficeToken(t),
				})

				return
			},
			expected: func(scenario fixtures.ScheduleDefaultOutput) schedule_gateway.PaginateData {
				return schedule_gateway.PaginateData{
					StartDate: scenario.StartDate,
					EndDate:   scenario.EndDate,
					Notes:     "",
					TutorID:   scenario.TutorID,
					PetIDs:    []string{scenario.PetID},
					Status:    schedule_status.Paid,
				}
			},
		},

		{
			title:          fmt.Sprintf("should not be able to patch status to %s if user role is %s", schedule_status.Paid, role.User),
			expectedStatus: http.StatusNotFound,
			seed: func() (response string, statusCode int, scenario fixtures.ScheduleDefaultOutput) {
				scenario = fixtures.Schedule.CreateDefault(t)

				statusCode = fixtures.Patch(t, fixtures.PatchInput{
					URI: fixtures.Schedule.URI + scenario.ID,
					Body: schedule_case.PatchValues{
						Status: schedule_status.Paid,
					},
					Response: &response,
					Token:    scenario.HostToken,
				})

				return
			},
			expected: func(scenario fixtures.ScheduleDefaultOutput) schedule_gateway.PaginateData {
				return schedule_gateway.PaginateData{
					StartDate: scenario.StartDate,
					EndDate:   scenario.EndDate,
					Notes:     "",
					TutorID:   scenario.TutorID,
					PetIDs:    []string{scenario.PetID},
					Status:    schedule_status.Open,
				}
			},
		},
	}

	for _, test := range tests {
		t.Run(test.title, func(t *testing.T) {
			fixtures.CleanDatabase()
			patchResponse, statusCode, scenario := test.seed()
			require.Equal(t, test.expectedStatus, statusCode, fmt.Sprintf("response: %s", patchResponse))

			page := schedule_gateway.PaginateOutput{
				Data: []schedule_gateway.PaginateData{},
			}

			fixtures.Get(t, fixtures.GetInput{
				URI:      fixtures.Schedule.URI,
				Response: &page,
				Token:    scenario.HostToken,
			})

			expected := test.expected(scenario)
			if len(page.Data) == 0 {
				require.Equal(t, expected, schedule_gateway.PaginateData{})
			} else {
				found := page.Data[0]
				require.Equal(t, expected.StartDate.Unix(), found.StartDate.Unix())
				require.Equal(t, expected.EndDate.Unix(), found.EndDate.Unix())
				expected.EndDate = time.Time{}
				expected.StartDate = time.Time{}
				found.StartDate = time.Time{}
				found.EndDate = time.Time{}

				require.Equal(t, expected, found)
			}
		})
	}
}
