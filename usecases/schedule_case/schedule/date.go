package schedule

import "time"

type DaysOfMonth uint32

type DateMonthYear struct {
	DaysOfMonth DaysOfMonth
	MonthYear   time.Time
}

func ToSchedule(startDate time.Time, endDate time.Time) []DateMonthYear {
	isSameMonthYear := startDate.Month() == endDate.Month() && startDate.Year() == endDate.Year()
	if isSameMonthYear {
		daysBetween := endDate.Day() - startDate.Day() + 1
		daysBetweenToBits := ((1 << daysBetween) - 1) << (startDate.Day() - 1)
		return []DateMonthYear{
			{
				MonthYear:   startDate,
				DaysOfMonth: DaysOfMonth(daysBetweenToBits),
			},
		}
	}

	dates := []DateMonthYear{}
	for month := startDate.Month(); month <= endDate.Month() || startDate.Year() < endDate.Year(); month++ {
		lastDayOfMonth := time.Date(startDate.Year(), month+1, 0, 0, 0, 0, 0, time.UTC)
		if month == endDate.Month() {
			lastDayOfMonth = endDate
		}
		daysBetween := lastDayOfMonth.Day() - startDate.Day() + 1
		daysBetweenToBits := ((1 << daysBetween) - 1) << (startDate.Day() - 1)
		dates = append(dates, DateMonthYear{
			MonthYear:   startDate,
			DaysOfMonth: DaysOfMonth(daysBetweenToBits),
		})
		startDate = lastDayOfMonth.AddDate(0, 0, 1)
	}

	return dates
}
