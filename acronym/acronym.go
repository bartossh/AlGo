// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package acronym should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package acronym

import (
	"strings"
)

func isSeparator(r rune) bool {
	switch r {
	case '_':
		return true
	case '-':
		return true
	case ' ':
		return true
	default:
		return false
	}
}

// Abbreviate should have a comment documenting it.
func Abbreviate(s string) string {
	// Write some code here to pass the test suite.
	// Then remove all the stock comments.
	// They're here to help you get started but they only clutter a finished solution.
	// If you leave them in, reviewers may protest!
	result := ""
	wasSeparator := true
	for _, l := range s {
		if wasSeparator && !isSeparator(l) {
			result += string(l)
		}
		wasSeparator = isSeparator(l)
	}
	return strings.ToUpper(result)
}
