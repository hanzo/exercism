// Package scrabble contains functions for scoring Scrabble words
package scrabble

import "strings"

var LetterScores = map[rune]int{
	'A': 1,
	'E': 1,
	'I': 1,
	'O': 1,
	'U': 1,
	'L': 1,
	'N': 1,
	'R': 1,
	'S': 1,
	'T': 1,
	'D': 2,
	'G': 2,
	'B': 3,
	'C': 3,
	'M': 3,
	'P': 3,
	'F': 4,
	'H': 4,
	'V': 4,
	'W': 4,
	'Y': 4,
	'K': 5,
	'J': 8,
	'X': 8,
	'Q': 10,
	'Z': 10,
}

// Score returns the scrabble score for the given word.
func Score(word string) int {
	var score int
	word = strings.ToUpper(word)

	// note that `range word` is the equivalent of `range []rune(word)`
	// https://staticcheck.io/docs/checks#S1029
	for _, rune := range word {
		// this lookup will return 0 for keys that aren't found
		score += LetterScores[rune]
	}
	return score
}
