/*
Package triangle implements a method for classifying triangles
 */
package triangle

import "math"

type Kind uint

const (
    // Pick values for the following identifiers used by the test program.
    NaT  = iota // not a triangle
    Equ // equilateral
    Iso // isosceles
    Sca // scalene
)

// Determine if a triangle is equilateral, isosceles, or scalene.
func KindFromSides(a, b, c float64) Kind {
	if !isTriangle(a, b, c) {
		return NaT
	}

	if a == b && b == c {
		return Equ
	}

	if a == b || b == c || a == c {
		return Iso
	}

	if a != b && b != c && a != c {
		return Sca
	}
	return NaT
}

func isTriangle(a, b, c float64) bool {
	return !math.IsInf(a, 0) && !math.IsInf(b, 0) && !math.IsInf(c, 0) && (a > 0 && b > 0 && c > 0) && (a + b >= c && a + c >= b && b + c >= a)
}