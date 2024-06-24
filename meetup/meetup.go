package meetup

import (
	"time"
)

// Define the WeekSchedule type here.
type WeekSchedule string

const (
	First  WeekSchedule = "first"
	Second WeekSchedule = "second"
	Third  WeekSchedule = "third"
	Fourth WeekSchedule = "fourth"
	Teenth WeekSchedule = "teenth"
	Last   WeekSchedule = "last"
)

func Day(wSched WeekSchedule, wDay time.Weekday, month time.Month, year int) int {
	first := time.Date(year, month, 1, 0, 0, 0, 0, time.UTC)

	for first.Weekday() != wDay {
		first = first.Add(time.Hour * 24)
	}

	if wSched == First {
		return first.Day()
	}

	second := first.Add(time.Hour * 24 * 7)
	if wSched == Second || wSched == Teenth && second.Day() >= 13 && second.Day() <= 19 {
		return second.Day()
	}

	third := second.Add(time.Hour * 24 * 7)
	if wSched == Third || wSched == Teenth && third.Day() >= 13 && third.Day() <= 19 {
		return third.Day()
	}

	fourth := third.Add(time.Hour * 24 * 7)
	if wSched == Fourth {
		return fourth.Day()
	}

	fifth := fourth.Add(time.Hour * 24 * 7)
	if wSched == Last {
		if fifth.Month() == month {
			return fifth.Day()
		}
		return fourth.Day()
	}

	return 0
}
