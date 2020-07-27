package clock

import (
	"fmt"
)

// Clock handles times without dates.
type Clock struct {
	// minuteOfDay is an integer between 0 and 1439 (inclusive) representing
	// the minute of the day (e.g. 1439 represents 23:59)
	minuteOfDay int
}

// New returns a Clock storing the time for the given hour and minute values.
func New(hour, minute int) Clock {
	m := minute + hour*60
	m %= 24 * 60
	// negative times should roll over to 'previous day'
	if m < 0 {
		m += 24 * 60
	}
	return Clock{m}
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d",
		c.minuteOfDay/60,
		c.minuteOfDay%60)
}

func (c Clock) Add(minutes int) Clock {
	return New(0, c.minuteOfDay+minutes)
}

func (c Clock) Subtract(minutes int) Clock {
	return New(0, c.minuteOfDay-minutes)
}
