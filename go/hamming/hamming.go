// Package hamming contains functions to calculate Hamming distance.
package hamming

import (
	"errors"
)

// Distance returns the Hamming distance between two DNA strands, or returns
// an error if the strands don't have the same length.
func Distance(a, b string) (int, error) {
	aRunes := []rune(a)
	bRunes := []rune(b)
	if len(aRunes) != len(bRunes) {
		return 0, errors.New("both DNA strands must have the same length")
	}
	distance := 0
	for i := range aRunes {
		if aRunes[i] != bRunes[i] {
			distance++
		}
	}
	return distance, nil
}
