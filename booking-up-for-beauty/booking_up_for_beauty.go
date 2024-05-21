package booking

import (
	"fmt"
	"time"
)

var opt int = -1

var layout map[int]string = map[int]string{
	1: "1/2/2006 15:04:05",
	2: "January 2, 2006 15:04:05",
	3: "Monday, January 2, 2006 15:04:05",
	4: "Monday, January 2, 2006, at 15:04.",
}

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
	if opt == -1 {
		opt = 1
	}
	v, _ := time.Parse(layout[opt], date)
	opt = -1
	return v
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	opt = 2
	return Schedule(date).Before(time.Now())
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	opt = 3
	hour := Schedule(date).Hour()
	return hour >= 12 && hour < 18
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	opt = 1
	return fmt.Sprintf("You have an appointment on %v", Schedule(date).Format(layout[4]))
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	return time.Date(time.Now().Year(), time.September, 15, 0, 0, 0, 0, time.UTC)
}
