// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package bob should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package bob

import (
	"fmt"
	"strings"
	"unicode"
)

func isUpper(s string) bool {
	for _, r := range s {
		if !unicode.IsUpper(r) && unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

func isQuestion(s string) bool {
	var c []string = strings.Split(s, "")

	for i, j := 0, len(c)-1; i < j; i, j = i+1, j-1 {
		c[i], c[j] = c[j], c[i]
	}

	reversed := strings.Join(c, "")
	for _, l := range string(reversed) {
		if byte(l) != byte(' ') && byte(l) != byte('\t') {
			if byte(l) == byte('?') {
				return true
			}
			return false
		} else {
			continue
		}
	}
	return false
}

func isLetter(s string) bool {
	for _, r := range s {
		if unicode.IsLetter(r) {
			fmt.Printf("%v\n", string(r))
			return true
		}
	}
	return false
}

func isEmpty(s string) bool {
	for _, l := range s {
		if byte(l) != byte(' ') && byte(l) != byte('\t') && byte(l) != byte('\r') && byte(l) != byte('\n') {
			return false
		}
	}
	return true
}

// Hey should have a comment documenting it.
func Hey(remark string) string {
	// Write some code here to pass the test suite.
	// Then remove all the stock comments.
	// They're here to help you get started but they only clutter a finished solution.
	// If you leave them in, reviewers may protest!

	if len(remark) == 0 || isEmpty(remark) {
		return "Fine. Be that way!"
	}

	if isUpper(remark) && isQuestion(remark) && isLetter(remark) {
		return "Calm down, I know what I'm doing!"
	}

	if isUpper(remark) && isLetter(remark) {
		return "Whoa, chill out!"
	}

	if isQuestion(remark) {
		return "Sure."
	}

	return "Whatever."
}
