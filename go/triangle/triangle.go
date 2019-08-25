// Package triangle contains functions to determine the Kind of a triangle.
package triangle

import "math"

type Kind int

const (
	NaT Kind = iota // not a triangle
	Equ             // equilateral
	Iso             // isosceles
	Sca             // scalene
)

// KindFromSides returns the kind of triangle that results from
// the three given side lengths.
func KindFromSides(a, b, c float64) Kind {
	var k Kind
	switch {
	case math.IsNaN(a) || math.IsNaN(b) || math.IsNaN(c):
		k = NaT
	case math.IsInf(a, 0) || math.IsInf(b, 0) || math.IsInf(c, 0):
		k = NaT
	case a < 0 || b < 0 || c < 0:
		k = NaT
	case a == 0 && b == 0 && c == 0:
		k = NaT
	case a+b < c || b+c < a || a+c < b:
		k = NaT
	case a == b && b == c:
		k = Equ
	case a == b || b == c || a == c:
		k = Iso
	default:
		k = Sca
	}
	return k
}
