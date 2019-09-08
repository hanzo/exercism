// Package accumulate contains functions for converting string lists
package accumulate

// Accumulate applies the given conversion to the list of strings and returns the modified strings
func Accumulate(input []string, convert func(string) string) []string {
	for i := range input {
		input[i] = convert(input[i])
	}
	return input
}
