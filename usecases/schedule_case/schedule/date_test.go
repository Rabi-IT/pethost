package schedule_test

import (
	"fmt"
	"pethost/usecases/schedule_case/schedule"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func Test_ToSchedule(t *testing.T) {
	type test struct {
		title     string
		startDate time.Time
		endDate   time.Time
		expected  []schedule.DateMonthYear
	}

	tests := []test{
		{
			title:     "same month",
			startDate: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
			endDate:   time.Date(2021, 1, 2, 0, 0, 0, 0, time.UTC),
			expected: []schedule.DateMonthYear{
				{
					MonthYear:   time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
					DaysOfMonth: 3, // 11 (2 days)
				},
			},
		},
		{
			title:     "same month",
			startDate: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
			endDate:   time.Date(2021, 1, 3, 0, 0, 0, 0, time.UTC),
			expected: []schedule.DateMonthYear{
				{
					MonthYear:   time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
					DaysOfMonth: 7, // 111 (3 days)
				},
			},
		},
		{
			title:     "same month",
			startDate: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
			endDate:   time.Date(2021, 1, 4, 0, 0, 0, 0, time.UTC),
			expected: []schedule.DateMonthYear{
				{
					MonthYear:   time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
					DaysOfMonth: 15, // 1111 (4 days)
				},
			},
		},
		{
			title:     "same month",
			startDate: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
			endDate:   time.Date(2021, 1, 5, 0, 0, 0, 0, time.UTC),
			expected: []schedule.DateMonthYear{
				{
					MonthYear:   time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
					DaysOfMonth: 31, // 11111 (5 days)
				},
			},
		},
		{
			title:     "same month",
			startDate: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
			endDate:   time.Date(2021, 1, 6, 0, 0, 0, 0, time.UTC),
			expected: []schedule.DateMonthYear{
				{
					MonthYear:   time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
					DaysOfMonth: 63, // 111111 (6 days)
				},
			},
		},
		{
			title:     "same month",
			startDate: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
			endDate:   time.Date(2021, 1, 7, 0, 0, 0, 0, time.UTC),
			expected: []schedule.DateMonthYear{
				{
					MonthYear:   time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
					DaysOfMonth: 127, // 1111111 (7 days)
				},
			},
		},
		{
			title:     "same month",
			startDate: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
			endDate:   time.Date(2021, 1, 8, 0, 0, 0, 0, time.UTC),
			expected: []schedule.DateMonthYear{
				{
					MonthYear:   time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
					DaysOfMonth: 255, // 11111111 (8 days)
				},
			},
		},
		{
			title:     "same month",
			startDate: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
			endDate:   time.Date(2021, 1, 9, 0, 0, 0, 0, time.UTC),
			expected: []schedule.DateMonthYear{
				{
					MonthYear:   time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
					DaysOfMonth: 511, // 111111111 (9 days)
				},
			},
		},
		{
			title:     "same month",
			startDate: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
			endDate:   time.Date(2021, 1, 10, 0, 0, 0, 0, time.UTC),
			expected: []schedule.DateMonthYear{
				{
					MonthYear:   time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
					DaysOfMonth: 1023, // 1111111111 (10 days)
				},
			},
		},
		{
			title:     "diff month",
			startDate: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
			endDate:   time.Date(2021, 2, 1, 0, 0, 0, 0, time.UTC),
			expected: []schedule.DateMonthYear{
				{
					MonthYear:   time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
					DaysOfMonth: (1 << 31) - 1, // 1111111111111111111111111111111 (31 days)
				},
				{
					MonthYear:   time.Date(2021, 2, 1, 0, 0, 0, 0, time.UTC),
					DaysOfMonth: 1, // 1 (1 day)
				},
			},
		},
		{
			title:     "diff month",
			startDate: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
			endDate:   time.Date(2021, 2, 2, 0, 0, 0, 0, time.UTC),
			expected: []schedule.DateMonthYear{
				{
					MonthYear:   time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
					DaysOfMonth: (1 << 31) - 1, // 1111111111111111111111111111111 (31 days)
				},
				{
					MonthYear:   time.Date(2021, 2, 1, 0, 0, 0, 0, time.UTC),
					DaysOfMonth: 3, // 11 (2 days)
				},
			},
		},
		{
			title:     "diff year",
			startDate: time.Date(2021, 12, 1, 0, 0, 0, 0, time.UTC),
			endDate:   time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC),
			expected: []schedule.DateMonthYear{
				{
					MonthYear:   time.Date(2021, 12, 1, 0, 0, 0, 0, time.UTC),
					DaysOfMonth: (1 << 31) - 1, // 1111111111111111111111111111111 (31 days)
				},
				{
					MonthYear:   time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC),
					DaysOfMonth: 1, // 1
				},
			},
		},

		{
			title:     "bisextile year",
			startDate: time.Date(2024, 2, 1, 0, 0, 0, 0, time.UTC),
			endDate:   time.Date(2024, 3, 1, 0, 0, 0, 0, time.UTC),
			expected: []schedule.DateMonthYear{
				{
					MonthYear:   time.Date(2024, 2, 1, 0, 0, 0, 0, time.UTC),
					DaysOfMonth: (1 << 29) - 1, // 11111111111111111111111111111 (29 days)
				},
				{
					MonthYear:   time.Date(2024, 3, 1, 0, 0, 0, 0, time.UTC),
					DaysOfMonth: 1, // 1
				},
			},
		},
	}

	for i, tc := range tests {
		title := fmt.Sprintf(
			"i: %d, startDate: %s, endDate: %s",
			i, tc.startDate.Format(time.RFC3339), tc.endDate.Format(time.RFC3339),
		)

		t.Run(title, func(t *testing.T) {
			result := schedule.ToSchedule(
				tc.startDate,
				tc.endDate,
			)

			for i, r := range result {
				require.Equal(t, tc.expected[i].MonthYear.Month(), r.MonthYear.Month())
				require.Equal(t, tc.expected[i].MonthYear.Year(), r.MonthYear.Year())
				require.Equal(t, tc.expected[i].DaysOfMonth, r.DaysOfMonth)
			}
		})
	}
}
