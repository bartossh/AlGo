package tries

import (
	"errors"
	"strings"
)

const (
	firstRune   = 97
	lengthRunes = 26
)

type FastNode struct {
	children [lengthRunes]*FastNode
	terminal bool
}

// NewFast creates FastNode entity and returns pointer to that entity
func NewFast() *FastNode {
	return &FastNode{}
}

// Insert inserts string to the FastNode tries structure if string is not numerical ASCII char
// or returns error otherwise
func (n *FastNode) Insert(s string) (bool, error) {
	s = strings.ToLower(s)
	for _, r := range s {
		if r < firstRune || r > firstRune+lengthRunes {
			return false, errors.New("rune out of scope, only non numerical ASCII chars are allowed")
		}
	}
	for _, r := range s {
		r = r - firstRune
		nn := n.children[r]
		if nn == nil {
			nn = NewFast()
			n.children[r] = nn
		}
		n = nn
	}
	if n.terminal {
		return false, nil
	}
	n.terminal = true
	return true, nil
}

// Find finds string in FastNode structure
func (n *FastNode) Find(s string) bool {
	s = strings.ToLower(s)
	for _, r := range s {
		if r < firstRune || r > firstRune+lengthRunes {
			return false
		}
		r = r - firstRune
		if nn := n.children[r]; nn != nil {
			n = nn
			continue
		}
		return false
	}
	return n.terminal
}

// Delete finds string in FastNode structure
func (n *FastNode) Delete(s string) bool {
	root := n
	s = strings.ToLower(s)
	for _, r := range s {
		if r < firstRune || r > firstRune+lengthRunes {
			return false
		}
		r = r - firstRune
		nn := n.children[r]
		if nn == nil {
			nn = NewFast()
			n.children[r] = nn
		}
		n = nn
	}
	if !n.terminal {
		return false
	}
	n.terminal = false

	root.clean()
	return true

}

func (n *FastNode) clean() bool {
	if n == nil {
		return false
	}

	if n.terminal {
		return true
	}

	var hasChilde bool
	for k, c := range n.children {
		if c.clean() {
			hasChilde = true
			continue
		}
		n.children[k] = nil
	}

	return hasChilde
}
