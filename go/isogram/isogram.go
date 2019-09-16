// Package isogram contains functions related to isograms
package isogram

import "unicode"

// IsIsogram returns true if the given word is an isogram (no repeated characters)
func IsIsogram(word string) bool {
	var seenLetters = make(map[rune]bool)
	for _, letter := range word {
		if unicode.IsSpace(letter) || letter == '-' {
			continue
		}
		letter = unicode.ToLower(letter)
		if _, found := seenLetters[letter]; found {
			return false
		} else {
			seenLetters[letter] = true
		}
	}
	return true
}
