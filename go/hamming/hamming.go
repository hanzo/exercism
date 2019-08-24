// Package hamming contains functions to calculate Hamming distance.
package hamming

import "errors"

// Distance returns the Hamming distance between two DNA strands, or returns
// an error if the strands don't have the same length.
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return -1, errors.New("Both DNA strands must have the same length")
	}
	distance := 0
	for i := range a {
		if a[i] != b[i] {
			distance++
		}
	}
	return distance, nil
}
