// Package strand contains functions for transcribing DNA and RNA strands.
package strand

import "strings"

var NucleotideComplements = map[rune](rune){
	'G': 'C',
	'C': 'G',
	'T': 'A',
	'A': 'U',
}

// Given a DNA strand, return its RNA complement (per RNA transcription).
func ToRNA(dna string) string {
	var builder strings.Builder
	for _, nucleotide := range dna {
		if complement, found := NucleotideComplements[nucleotide]; found {
			builder.WriteRune(complement)
		} else {
			builder.WriteRune(nucleotide)
		}
	}
	return builder.String()
}
