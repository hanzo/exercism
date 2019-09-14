// Package scrabble contains functions for scoring Scrabble words
package scrabble

import "unicode"

// ScoreLetter returns the scrabble score for the given letter.
// Returns zero for non-English letters.
func ScoreLetter(letter rune) int {
	switch unicode.ToUpper(letter) {
	case 'A', 'E', 'I', 'O', 'U', 'L', 'N', 'R', 'S', 'T':
		return 1
	case 'D', 'G':
		return 2
	case 'B', 'C', 'M', 'P':
		return 3
	case 'F', 'H', 'V', 'W', 'Y':
		return 4
	case 'K':
		return 5
	case 'J', 'X':
		return 8
	case 'Q', 'Z':
		return 10
	default:
		return 0
	}
}

// Score returns the scrabble score for the given word.
func Score(word string) int {
	var score int
	// note that `range word` is the equivalent of `range []rune(word)`
	// https://staticcheck.io/docs/checks#S1029
	for _, rune := range word {
		score += ScoreLetter(rune)
	}
	return score
}
