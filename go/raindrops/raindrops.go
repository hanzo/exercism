// Package raindrops contains functions to convert numbers to raindrops strings.
package raindrops

import "strconv"

// Convert returns the raindrop string of the given integer.
func Convert(input int) string {
	var output string
	if input%3 == 0 {
		output += "Pling"
	}
	if input%5 == 0 {
		output += "Plang"
	}
	if input%7 == 0 {
		output += "Plong"
	}
	if output == "" {
		output = strconv.Itoa(input)
	}
	return output
}
