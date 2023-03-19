// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package leap should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package leap

// IsLeapYear should have a comment documenting it.
func IsLeapYear(year int) bool {
	devisableBy4 := year%4 == 0
	devisableBy100 := year%100 == 0
	devisableBy400 := year%400 == 0
	return devisableBy4 && (!devisableBy100 || devisableBy400)
}
