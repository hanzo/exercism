package clock

import (
	"fmt"
)

const (
	minutesPerHour = 60
	hoursPerDay    = 24
	minutesPerDay  = minutesPerHour * hoursPerDay
)

// Clock handles times without dates.
type Clock interface {
	String() string
	Add(minutes int) Clock
	Subtract(minutes int) Clock
}

type clock struct {
	// minuteOfDay is an integer between 0 and 1439 (inclusive) representing
	// the minute of the day (e.g. 1439 represents 23:59)
	minuteOfDay int
}

// New returns a Clock storing the time for the given hour and minute values.
func New(hour, minute int) Clock {
	totalMins := (hour * minutesPerHour) + minute
	trimmedMins := totalMins % minutesPerDay
	// negative times should roll over to 'previous day'
	if trimmedMins < 0 {
		trimmedMins += minutesPerDay
	}
	return clock{
		minuteOfDay: trimmedMins,
	}
}

func (c clock) String() string {
	return fmt.Sprintf("%02d:%02d",
		c.minuteOfDay/minutesPerHour,
		c.minuteOfDay%minutesPerHour)
}

func (c clock) Add(minutes int) Clock {
	return New(0, c.minuteOfDay+minutes)
}

func (c clock) Subtract(minutes int) Clock {
	return New(0, c.minuteOfDay-minutes)
}
