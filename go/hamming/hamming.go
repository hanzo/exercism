// Package hamming contains functions to calculate Hamming distance.
package hamming

import (
	"errors"
	"unicode/utf8"
)

// Distance returns the Hamming distance between two DNA strands, or returns
// an error if the strands don't have the same length.
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New("both DNA strands must have the same length")
	}
	distance := 0

	// Decode one rune at a time from each input string and compare their values and widths.
	// See https://blog.golang.org/strings for more details
	for i, w := 0, 0; i < len(a); i += w {
		aVal, aWidth := utf8.DecodeRuneInString(a[i:])
		bVal, bWidth := utf8.DecodeRuneInString(b[i:])
		if aVal != bVal || aWidth != bWidth {
			distance++
		}
		w = aWidth
	}

	return distance, nil
}
