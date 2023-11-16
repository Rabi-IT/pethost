package pet

import "time"

const (
	PuppieAge  = 1
	ElderlyAge = 10
)

func CalculateAge(birthdate, today time.Time) int {
	location := birthdate.Location()
	if location == nil {
		location = time.UTC
	}
	today = today.In(location)
	ty, tm, td := today.Date()
	today = time.Date(ty, tm, td, 0, 0, 0, 0, time.UTC)
	by, bm, bd := birthdate.Date()
	birthdate = time.Date(by, bm, bd, 0, 0, 0, 0, time.UTC)
	if today.Before(birthdate) {
		return 0
	}

	age := ty - by
	anniversary := birthdate.AddDate(age, 0, 0)
	if anniversary.After(today) {
		age--
	}
	return age
}
