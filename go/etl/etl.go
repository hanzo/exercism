// Package etl contains functions for extract/transform/load jobs
package etl

import "strings"

// Transform receives a mapping of scrabble scores to letters and transforms it
// into a mapping of lowercase letters to scores
func Transform(input map[int][]string) map[string]int {
	var output = make(map[string]int)
	for score, letters := range input {
		for _, letter := range letters {
			output[strings.ToLower(letter)] = score
		}
	}
	return output
}
