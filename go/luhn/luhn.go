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
	// double every second digit starting from the right
	doubleDigit := len(input)%2 == 0
	sum := 0
	for _, ch := range input {
		digit, err := strconv.Atoi(string(ch))
		if err != nil {
			return false
		}
		if doubleDigit {
			digit *= 2
			if digit > 9 {
				digit -= 9
			}
		}
		sum += digit
		doubleDigit = !doubleDigit
	}
	return sum%10 == 0
}
