// Package gigasecond contains time calculation functions
package gigasecond

import "time"

const Gigasecond time.Duration = time.Second * 1e9

// AddGigasecond returns the time one gigasecond later than the given time
func AddGigasecond(t time.Time) time.Time {
	return t.Add(Gigasecond)
}
