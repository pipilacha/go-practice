// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package triangle should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package triangle

type Kind string

const (
	NaT Kind = "NaT"
	Equ Kind = "Equ"
	Iso Kind = "Iso"
	Sca Kind = "Sca"
)

// KindFromSides return thes kind of a triangle based on the measurement of each side
// possible  values are :
//
//	NaT => not a triangle
//	Equ => equilateral
//	Iso => isosceles
//	Sca => scalene
func KindFromSides(a, b, c float64) Kind {

	if a <= 0 || b <= 0 || c <= 0 || !(a+b >= c && a+c >= b && b+c >= a) {
		return NaT
	}

	if a == b && b == c {
		return Equ
	}

	if a == b || a == c || b == c {
		return Iso
	}

	return Sca
}
