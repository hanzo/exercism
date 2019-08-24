// Package leap contains functions related to leap year calculation.
package leap

// IsLeapYear returns true if the given year is a leap year.
func IsLeapYear(year int) bool {
	switch {
	case year%4 == 0 && year%400 == 0:
		return true
	case year%4 == 0 && year%100 != 0:
		return true
	default:
		return false
	}
}
