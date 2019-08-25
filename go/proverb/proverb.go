// Package proverb contains functions that create proverbs.
package proverb

import "fmt"

const VerseTemplate = "For want of a %s the %s was lost."
const EndTemplate = "And all for the want of a %s."

// Proverb returns a string containing a proverb generated
// from the given list of words.
func Proverb(rhyme []string) []string {
	proverb := []string{}
	for i := range rhyme {
		if i == len(rhyme)-1 {
			proverb = append(proverb, fmt.Sprintf(EndTemplate, rhyme[0]))
		} else {
			proverb = append(proverb, fmt.Sprintf(VerseTemplate, rhyme[i], rhyme[i+1]))
		}
	}
	return proverb
}
