// Package acronym contains functions related to acronyms
package acronym

import (
	"strings"
	"unicode"
)

// Abbreviate returns the acronym for the given phrase
func Abbreviate(s string) string {
	var builder strings.Builder
	words := strings.FieldsFunc(s, Split)

	for _, word := range words {
		for _, char := range word {
			// find the first letter in the word, and add it to the acronym as a capital letter
			if unicode.IsLetter(char) {
				builder.WriteRune(unicode.ToUpper(char))
				break
			}
		}
	}
	return builder.String()
}

// Split delimits words on either whitespace characters or hyphens
func Split(r rune) bool {
	return unicode.IsSpace(r) || r == '-'
}
