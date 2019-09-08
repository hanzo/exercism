// Package romannumerals contains functions for converting between decimal numbers and roman numerals
package romannumerals

import (
	"errors"
	"strings"
)

var DecimalToNumeral = []struct {
	integer int
	numeral string
}{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

// ToRomanNumeral returns the roman numeral representation of the given decimal number
func ToRomanNumeral(decimal int) (string, error) {
	if decimal < 1 || decimal > 3000 {
		return "", errors.New("Input must be greater than 0 and less than 3001")
	}
	var builder strings.Builder
	for _, pair := range DecimalToNumeral {
		for decimal >= pair.integer {
			builder.WriteString(pair.numeral)
			decimal -= pair.integer
		}
	}
	return builder.String(), nil
}
