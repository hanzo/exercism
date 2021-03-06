// Package dna contains functions for counting nucleotides in DNA strands
package dna

import (
	"fmt"
)

// Histogram is a mapping from nucleotide to its count in given DNA.
type Histogram map[rune]int

// DNA is a list of nucleotides
type DNA string

// Counts generates a histogram of valid nucleotides in the given DNA.
// Returns an error if d contains an invalid nucleotide.
func (d DNA) Counts() (Histogram, error) {

	// initialize the histogram with a count of zero for each nucleotide type
	h := Histogram{
		'A': 0,
		'C': 0,
		'G': 0,
		'T': 0,
	}

	for _, nuc := range d {
		if _, found := h[nuc]; found {
			h[nuc]++
		} else {
			return nil, fmt.Errorf("found invalid nucleotide: %v", nuc)
		}
	}
	return h, nil
}
