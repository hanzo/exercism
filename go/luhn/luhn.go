// Package luhn contains functions related to the Luhn algorithm
package luhn

import (
	"strconv"
	"strings"
)

// Valid returns true if the given string contains a number
// which is valid according to the Luhn algorithm
func Valid(input string) bool {
	input = strings.ReplaceAll(input, " ", "")
	if len(input) <= 1 {
		return false
	}
	sum := 0
	shouldDouble := false
	for i := len(input) - 1; i >= 0; i-- {
		digit, err := strconv.Atoi(string(input[i]))
		if err != nil {
			return false
		}
		if shouldDouble {
			digit *= 2
			if digit > 9 {
				digit -= 9
			}
		}
		sum += digit
		shouldDouble = !shouldDouble
	}
	return sum%10 == 0
}
