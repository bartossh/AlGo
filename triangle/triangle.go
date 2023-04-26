// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package triangle should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package triangle

import (
	"math"
	"sort"
)

// Notice KindFromSides() returns this type. Pick a suitable data type.
type Kind int

const (
	// Pick values for the following identifiers used by the test program.
	NaT Kind = iota // not a triangle
	Equ             // equilateral
	Iso             // isosceles
	Sca             // scalene
)

// KindFromSides should have a comment documenting it.
func KindFromSides(a, b, c float64) Kind {
	sides := []float64{a, b, c}
	sort.Float64s(sides)
	if sides[0]+sides[1] <= sides[2] {
		return NaT
	}
	sideLength := make(map[float64]int)
	for _, v := range sides {
		if v <= 0.0 || math.IsNaN(v) {
			return NaT
		}
		sideLength[v]++
	}
	switch len(sideLength) {
	case 1:
		return Equ
	case 2:
		return Iso
	default:
		return Sca
	}
}
