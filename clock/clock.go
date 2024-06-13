package clock

import (
	"fmt"
)

// Define the Clock type here.

type Clock struct {
	hours   int
	minutes int
}

func calcClock(h int, m int) (hrs int, mins int) {

	total := h*60 + m

	if total%60 != 0 {
		mins = total % 60
		if total < 0 {
			mins += 60
		}

		total -= mins
	}

	hrs = total / 60 % 24

	if total < 0 && hrs != 0 {
		hrs += 24
	}

	return hrs, mins
}

func New(h, m int) Clock {
	hrs, mins := calcClock(h, m)
	return Clock{hours: hrs, minutes: mins}
}

func (c Clock) Add(m int) Clock {
	c.hours, c.minutes = calcClock(c.hours, c.minutes+m)
	return c
}

func (c Clock) Subtract(m int) Clock {
	c.hours, c.minutes = calcClock(c.hours, c.minutes-m)
	return c
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.hours, c.minutes)
}
