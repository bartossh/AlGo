package tries

import "unicode/utf8"

// Node represents node that can be terminal or otherwise can have children
type Node struct {
	children map[rune]*Node
	terminal bool
}

// New creates a new Node entity and returns pointer to that entity,
// Created node is not set to be termianl
func New() *Node {
	children := make(map[rune]*Node)
	return &Node{children, false}
}

// Insert inserts string in to the tries node structure
func (n *Node) Insert(s string) bool {
	for i, w := 0, 0; i < len(s); i += w {
		l, width := utf8.DecodeRuneInString(s[i:])
		w = width

		nf, ok := n.children[l]
		if ok {
			n = nf
			continue
		}
		nn := New()
		n.children[l] = nn
		n = nn
	}

	if n.terminal {
		return false
	}

	n.terminal = true
	return true
}

// Find finds node in the tries node structure
func (n *Node) Find(s string) bool {
	for i, w := 0, 0; i < len(s); i += w {
		l, width := utf8.DecodeRuneInString(s[i:])
		w = width
		nf, ok := n.children[l]
		if !ok {
			return false
		}
		n = nf
	}
	if n.terminal {
		return true
	}
	return false
}