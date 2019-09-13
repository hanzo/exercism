// Package bob contains functions related to Bob the teenager
package bob

import (
	"strings"
	"unicode"
)

// Hey returns Bob's response to the given remark
func Hey(remark string) string {
	trimmed := strings.TrimSpace(remark)
	if len(trimmed) == 0 {
		return "Fine. Be that way!"
	}

	isQuestion := strings.HasSuffix(trimmed, "?")
	hasLetters := false
	allCaps := true

	for _, char := range trimmed {
		if unicode.IsLetter(char) {
			hasLetters = true
			if unicode.IsLower(char) {
				allCaps = false
			}
		}
	}

	if hasLetters && allCaps && isQuestion {
		return "Calm down, I know what I'm doing!"
	} else if hasLetters && allCaps {
		return "Whoa, chill out!"
	} else if isQuestion {
		return "Sure."
	} else {
		return "Whatever."
	}
}
