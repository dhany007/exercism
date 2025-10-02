// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package triangle should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package triangle

// Notice KindFromSides() returns this type. Pick a suitable data type.
type Kind string

const (
	// Pick values for the following identifiers used by the test program.
	NaT Kind = "not a triangle" // not a triangle
	Equ Kind = "equilateral"
	Iso Kind = "isosceles"
	Sca Kind = "scalene"
)

// KindFromSides should have a comment documenting it.
func KindFromSides(a, b, c float64) Kind {
	var k Kind

	if !isTriangle(a, b, c) {
		return NaT
	}

	if a == b && b == c {
		return Equ
	}

	if a != b && b != c && a != c {
		return Sca
	}

	if a == b || b == c || c == a {
		return Iso
	}

	return k
}

func isTriangle(a, b, c float64) bool {
	if a == 0 && b == 0 && c == 0 {
		return false
	}

	exp1 := a+b >= c
	exp2 := b+c >= a
	exp3 := a+c >= b

	return exp1 && exp2 && exp3
}
