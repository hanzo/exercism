// Package grains contains functions for calculating the number of grains on a chessboard
package grains

import "errors"

// BoardTotal contains the sum of 2^1 + 2^2 + ... + 2^63
const BoardTotal = (1 << 64) - 1

// Total returns the total number of grains on a chessboard
func Total() uint64 {
	return BoardTotal
}

// Square returns the number of grains on the given chessboard square
func Square(squareNum int) (uint64, error) {
	if squareNum <= 0 || squareNum > 64 {
		return 0, errors.New("square number must be between 1 and 64 (inclusive)")
	}
	// Shift the binary number 1 to the left to create increasing powers of 2
	return 1 << (squareNum - 1), nil
}
