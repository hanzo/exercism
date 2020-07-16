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
	minuteOfDay int
}

// getMinuteOfDay returns an integer between 0 and 1439 (inclusive) representing
// the minute of the day for the given hour and minute values.
func getMinuteOfDay(hour, minute int) int {
	totalMins := (hour * minutesPerHour) + minute
	trimmedMins := totalMins % minutesPerDay
	// negative times should roll over to 'previous day'
	if trimmedMins < 0 {
		trimmedMins += minutesPerDay
	}
	return trimmedMins
}

// New returns a Clock storing the time for the given hour and minute values.
func New(hour, minute int) Clock {
	return clock{
		minuteOfDay: getMinuteOfDay(hour, minute),
	}
}

func (c clock) String() string {
	return fmt.Sprintf("%02d:%02d",
		c.minuteOfDay/minutesPerHour,
		c.minuteOfDay%minutesPerHour)
}

func (c clock) Add(minutes int) Clock {
	return clock{
		minuteOfDay: getMinuteOfDay(0, c.minuteOfDay+minutes),
	}
}

func (c clock) Subtract(minutes int) Clock {
	return clock{
		minuteOfDay: getMinuteOfDay(0, c.minuteOfDay-minutes),
	}
}
