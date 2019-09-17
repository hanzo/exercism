// Package diffsquares contains functions related for calculating the sum of squares and the square of sum
package diffsquares

import "math"

// SquareOfSum returns the square of the sum of the first N natural numbers
func SquareOfSum(n int) int {
	// Formula: https://en.wikipedia.org/wiki/1_%2B_2_%2B_3_%2B_4_%2B_%E2%8B%AF#Partial_sums
	sum := n * (n + 1) / 2
	return int(math.Pow(float64(sum), 2))
}

// SumOfSquares returns the sum of the square of each of the first N natural numbers
func SumOfSquares(n int) int {
	// Formula: https://www.math-only-math.com/sum-of-the-squares-of-first-n-natural-numbers.html
	return (n * (n + 1) * (2*n + 1)) / 6
}

// Difference returns the difference between the square of the sum and the sum of the squares of the first N natural numbers
func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
