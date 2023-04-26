// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package proverb should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package proverb

import "fmt"

// Proverb should have a comment documenting it.
func Proverb(rhyme []string) []string {
	if len(rhyme) == 0 {
		return []string{}
	}
	coll := make([]string, 0, len(rhyme))
	max := len(rhyme)
	for i := range rhyme {
		if i < max-1 {
			coll = append(coll, createLostMsg(rhyme[i], rhyme[i+1]))
		}
	}
	coll = append(coll, createWantMsg(rhyme[0]))

	return coll
}

func createLostMsg(first, second string) string {
	return fmt.Sprintf("For want of a %s the %s was lost.", first, second)
}

func createWantMsg(want string) string {
	return fmt.Sprintf("And all for the want of a %s.", want)
}
