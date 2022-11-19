package btree

import (
	"fmt"
	"strings"

	"golang.org/x/exp/constraints"
)

type comperable = constraints.Ordered

type btreeNode[T comperable] struct {
	parent *btreeNode[T]
	leafs  []*btreeNode[T]
	values []T
}

// Root is a top node in the btree data structure
type Root[T comperable] struct {
	root *btreeNode[T]
	max  int
}

func newBtreeNode[T comperable](max int, parent *btreeNode[T]) *btreeNode[T] {
	leafs := make([]*btreeNode[T], 0, max+1)
	values := make([]T, 0, max)
	return &btreeNode[T]{parent, leafs, values}
}

// New creates a new btree structure of given type T  returning pointer to the root node of that structure
func New[T comperable](max int) *Root[T] {
	n := newBtreeNode[T](max, nil)
	return &Root[T]{n, max}
}

func (n *btreeNode[T]) splitLeafsLastNode(v T, l, r []T) (nV T, nLN, nRN *btreeNode[T]) {
	ln := newBtreeNode[T](cap(n.values), n)
	rn := newBtreeNode[T](cap(n.values), n)
	ln.values = l
	rn.values = r
	return n.splitLeafsIntermediate(v, ln, rn)
}

func (n *btreeNode[T]) splitLeafsIntermediate(v T, ln, rn *btreeNode[T]) (nV T, nLN, nRN *btreeNode[T]) {
	ln.parent = n.parent
	rn.parent = n.parent
	switch {
	case len(n.leafs) == 0:
		n.leafs = append(n.leafs, []*btreeNode[T]{ln, rn}...)
		n.values = append(n.values, v)
		return
	case len(n.values) == cap(n.values):
		var tempLeafs []*btreeNode[T]
		var tempValues []T
		var f bool
	FoundValueLoop:
		for i, vv := range n.values {
			if v < vv {
				tempLeafs = append(n.leafs[:i], append([]*btreeNode[T]{ln, rn}, n.leafs[i+1:]...)...)
				tempValues = append(n.values[:i], append([]T{v}, n.values[i:]...)...)
				f = true
				break FoundValueLoop
			}
		}
		if !f {
			tempLeafs = append(n.leafs[:len(n.leafs)-1], []*btreeNode[T]{ln, rn}...)
			tempValues = append(n.values, v)
		}
		half := (len(tempValues) - 1) / 2
		nV = tempValues[half]
		n.values = make([]T, 0, cap(n.values))
		n.values = append(tempValues[:half], tempValues[half+1:]...)

		nLN = tempLeafs[half-1]
		nRN = tempLeafs[half]

		n.leafs = make([]*btreeNode[T], 0, cap(n.leafs))
		n.leafs = append(tempLeafs[:half], tempLeafs[half+2:]...)
		return
	default:
		for i, vv := range n.values {
			if v < vv {
				n.leafs = append(n.leafs[:i], append([]*btreeNode[T]{ln, rn}, n.leafs[i+1:]...)...)
				n.values = append(n.values[:i], append([]T{v}, n.values[i:]...)...)
				return
			}
		}
		n.leafs = append(n.leafs[:len(n.leafs)-1], []*btreeNode[T]{ln, rn}...)
		n.values = append(n.values, v)
		return
	}
}

func (n *btreeNode[T]) insert(v T) (T, []T, []T) {
	idx := len(n.values)
	for i, vv := range n.values {
		if v == vv {
			var t T
			return t, nil, nil
		}
		if v < vv {
			idx = i
			break
		}
	}
	return n.insertAt(idx, v)
}

func (n *btreeNode[T]) insertAt(i int, v T) (nv T, l, r []T) {
	temp := append(n.values[:i], append([]T{v}, n.values[i:]...)...)
	if len(n.values) == cap(n.values) {
		half := (len(temp) - 1) / 2
		nv = temp[half]
		l = make([]T, len(temp[:half]), cap(n.values))
		r = make([]T, len(temp[half+1:]), cap(n.values))
		copy(l, temp[:half])
		copy(r, temp[half+1:])
		return
	}
	n.values = temp
	return
}

// Insert inserts entity of type T in to the btree structure
func (r *Root[T]) Insert(v T) {
	node := r.root
FinderLoop:
	for {
		if len(node.leafs) == 0 {
			break FinderLoop
		}
		f := false
	ValuesLoop:
		for i, vv := range node.values {
			if v < vv {
				node = node.leafs[i]
				f = true
				break ValuesLoop
			}
		}
		if !f {
			node = node.leafs[len(node.leafs)-1]
		}
	}

	var left, right []T
	v, left, right = node.insert(v)
	if len(left) == 0 && len(right) == 0 {
		return
	}

	var leftNode, rightNode *btreeNode[T]
	parent := node.parent
	if parent == nil {
		parent = newBtreeNode[T](cap(node.values), nil)
		v, leftNode, rightNode = parent.splitLeafsLastNode(v, left, right)
		node.parent = parent
		r.root = parent
		node = parent
		return
	}

SpliterLoop:
	for {
		if parent == nil {
			parent = newBtreeNode[T](cap(node.values), nil)
			node.parent = parent
			r.root = parent
		}
		v, leftNode, rightNode = parent.splitLeafsIntermediate(v, leftNode, rightNode)
		if leftNode == nil && rightNode == nil {
			break SpliterLoop
		}
		node = parent
		parent = node.parent
	}
}

// Traversal travers the nodes printing the values
func (r *Root[T]) Traversal() {
	r.root.traversal(0)
}

func (n *btreeNode[T]) traversal(spcs int) {
	indent := strings.Repeat("-", spcs)
	for i := 0; i < len(n.values); i++ {
		if len(n.leafs) != 0 {
			n.leafs[i].traversal(spcs + 2)
		}
		fmt.Printf("%s%v\n", indent, n.values[i])
	}
	if len(n.leafs) != 0 {
		n.leafs[len(n.leafs)-1].traversal(spcs + 2)
	}
}

// Find looks for entity of type T in the btree structure
// returning true if entity exists or false otherwise
func (r *Root[T]) Find(v T) bool {
	node := r.root
NodeTraversal:
	for {
		for i, vv := range node.values {
			switch {
			case v == vv:
				return true
			case v < vv:
				if len(node.leafs) == 0 {
					break NodeTraversal
				}
				node = node.leafs[i]
				continue NodeTraversal
			}
		}
		if len(node.leafs) == 0 {
			break NodeTraversal
		}
		node = node.leafs[len(node.leafs)-1]
	}
	return false
}
