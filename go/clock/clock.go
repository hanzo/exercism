package clock

import (
	"fmt"
	"time"
)

// Clock handles times without dates.
type Clock interface {
	String() string
	Add(minutes int) Clock
	Subtract(minutes int) Clock
}

// Ideally we could store a time.Time on the struct instead of having to call time.Date()
// repeatedly, but this causes equality checks to fail in some cases (e.g. day wraparound).
type clock struct {
	hour   int
	minute int
}

// makeTime returns a Time object for the given hour and minute. The resulting
// Time object will have Hour and Minute values in their expected ranges.
// The other values (day, second, etc) are unused.
func makeTime(hour, minute int) time.Time {
	return time.Date(0, 0, 0, hour, minute, 0, 0, time.UTC)
}

// New returns a Clock with the given hour and minute values.
func New(hour, minute int) Clock {
	theTime := makeTime(hour, minute)
	return clock{
		hour:   theTime.Hour(),
		minute: theTime.Minute(),
	}
}

func (c clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.hour, c.minute)
}

func (c clock) Add(minutes int) Clock {
	oldTime := makeTime(c.hour, c.minute)
	newTime := oldTime.Add(time.Minute * time.Duration(minutes))
	return clock{
		minute: newTime.Minute(),
		hour:   newTime.Hour(),
	}
}

func (c clock) Subtract(minutes int) Clock {
	oldTime := makeTime(c.hour, c.minute)
	newTime := oldTime.Add(time.Minute * time.Duration(minutes) * -1)
	return clock{
		minute: newTime.Minute(),
		hour:   newTime.Hour(),
	}
}
