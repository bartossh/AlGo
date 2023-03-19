package linkedlist

import (
	"errors"
)

// ErrorEmptyList error of empty list
type ErrorEmptyList error

// NewErrEmptyList returns error of empty list
func NewErrEmptyList() ErrorEmptyList {
	return errors.New("cannot pop from list of zero length")
}

// ErrEmptyList error representation of ErrorEmpty List
var ErrEmptyList = NewErrEmptyList()

// Node represents double linked list node
type Node struct {
	first, last, next, prev *Node
	Val                     interface{}
}

// NodeActions provides interface to move between Nodes
type NodeActions interface {
	Next() NodeActions
	Prev() NodeActions
	First() NodeActions
	Last() NodeActions
}

// Next returns pointer to next Node
func (n *Node) Next() *Node {
	return n.next
}

// Prev returns pointer to previous Node
func (n *Node) Prev() *Node {
	return n.prev
}

// First returns pointer to previous Node
func (n *Node) First() *Node {
	return n.first
}

// Last returns pointer to previous Node
func (n *Node) Last() *Node {
	return n.last
}

// List represents double linked list
type List struct {
	list []*Node
}

// ListActions provides interface to access or manipulate List values
type ListActions interface {
	First() NodeActions
	Last() NodeActions
	PushBack(v interface{})
	PopBack() (interface{}, error)
	PushFront(v interface{})
	PopFront() (interface{}, error)
	Reverse()
}

// NewList returns pointer to newly populated with args double linked List
func NewList(args ...interface{}) *List {
	l := List{}
	if len(args) == 0 {
		return &l
	}
	list := make([]*Node, len(args))
	for i, v := range args {
		var n Node
		n.Val = v
		list[i] = &n
	}
	for i, n := range list {
		if i == 0 {
			n.prev = nil
		} else {
			n.prev = list[i-1]
		}
		n.first = list[0]
		if i < len(list)-1 {
			n.next = list[i+1]
		} else {
			n.next = nil
		}
		n.last = list[len(list)-1]
	}
	l.list = list
	return &l
}

// First returns first node in double linked List
func (l *List) First() *Node {
	if len(l.list) == 0 {
		return nil
	}
	return l.list[0].First()
}

// Last returns last node in double linked List
func (l *List) Last() *Node {
	if len(l.list) == 0 {
		return nil
	}
	return l.list[len(l.list)-1].Last()
}

// PushBack pushes new Node on the back of double linked List
func (l *List) PushBack(v interface{}) {
	node := &Node{}
	node.Val = v
	l.list = append(l.list, node)
	l.list[len(l.list)-1].first = l.list[0]
	for _, v := range l.list {
		v.last = l.list[len(l.list)-1]
	}
	if len(l.list) > 1 {
		l.list[len(l.list)-1].prev = l.list[len(l.list)-2]
		l.list[len(l.list)-2].next = l.list[len(l.list)-1]
	} else {
		l.list[len(l.list)-1].prev = nil
	}
	l.list[len(l.list)-1].next = nil
}

// PopBack pops value from the last node in double linked List
func (l *List) PopBack() (interface{}, error) {
	if len(l.list) == 0 {
		return nil, ErrEmptyList
	}
	v := l.list[len(l.list)-1]
	if len(l.list) == 1 {
		l.list = make([]*Node, 0)
		return v.Val, nil
	}
	l.list = append(l.list[:len(l.list)-1])
	l.list[len(l.list)-1].next = nil
	for _, v := range l.list {
		v.last = l.list[len(l.list)-1]
	}
	return v.Val, nil
}

// PushFront pushes new Node on the front of double linked List
func (l *List) PushFront(v interface{}) {
	n := &Node{Val: v}
	l.list = append([]*Node{n}, l.list...)
	l.list[0].last = l.list[len(l.list)-1]
	if len(l.list) > 1 {
		l.list[0].next = l.list[1]
		l.list[1].prev = l.list[0]
	} else {
		l.list[0].next = nil
	}
	for _, v := range l.list {
		v.first = l.list[0]
	}
}

// PopFront pops value from the first position in double linked List
func (l *List) PopFront() (interface{}, error) {
	if len(l.list) == 0 {
		return nil, ErrEmptyList
	}
	v := l.list[0]
	if len(l.list) == 1 {
		l.list = make([]*Node, 0)
		return v.Val, nil
	}
	l.list = l.list[1:]
	l.list[0].first = nil
	l.list[0].prev = nil
	for _, v := range l.list {
		v.first = l.list[0]
	}
	return v.Val, nil
}

// Reverse reverses double linked list
func (l *List) Reverse() {
	for i, v := range l.list {
		v.prev, v.next = v.next, v.prev
		v.last = l.list[0]
		v.first = l.list[len(l.list)-1]
		if i == 0 {
			v.next = nil
		}
		if i == len(l.list)-1 {
			v.prev = nil
		}
	}
}
