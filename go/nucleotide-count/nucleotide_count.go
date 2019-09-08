// Package dna contains functions for counting nucleotides in DNA strands
package dna

import (
	"fmt"
	"strings"
)

// Histogram is a mapping from nucleotide to its count in given DNA.
type Histogram map[rune]int

// DNA is a list of nucleotides
type DNA string

// Nucleotides is a static list of valid nucleotide types
var Nucleotides = [4]rune{'A', 'C', 'G', 'T'}

// Counts generates a histogram of valid nucleotides in the given DNA.
// Returns an error if d contains an invalid nucleotide.
func (d DNA) Counts() (Histogram, error) {

	// initialize the histogram with a count of zero for each nucleotide type
	var h Histogram = make(Histogram)
	for _, nuc := range Nucleotides {
		h[nuc] = 0
	}

	dna := strings.ToUpper(string(d))
	for _, nuc := range dna {
		if _, found := h[nuc]; found {
			h[nuc]++
		} else {
			return nil, fmt.Errorf("found invalid nucleotide: %v", nuc)
		}
	}
	return h, nil
}
