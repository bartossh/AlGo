package btree

import (
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

func (n *btreeNode[T]) splitLeafs(v T, l []T, r []T) {
	ln := newBtreeNode[T](cap(n.values), n)
	rn := newBtreeNode[T](cap(n.values), n)
	ln.values = l
	rn.values = r
	if len(n.leafs) == 0 {
		n.values = append(n.values, v)
		n.leafs = append(n.leafs, []*btreeNode[T]{ln, rn}...)
		return
	}
	for i, vv := range n.values {
		if v < vv {
			n.leafs = append(n.leafs[:i], append([]*btreeNode[T]{ln, rn}, n.leafs[i+1:]...)...)
			n.values = append(n.values[:i], append([]T{v}, n.values[i:]...)...)
			return
		}
	}
	n.leafs = append(n.leafs[:len(n.leafs)-1], []*btreeNode[T]{ln, rn}...)
	n.values = append(n.values, v)
}

func (n *btreeNode[T]) insert(v T) (T, []T, []T) {
	idx := 0
	for i, vv := range n.values {
		if v == vv {
			var t T
			return t, nil, nil
		}
		if v > vv {
			idx = i + 1
			continue
		}
		break
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
		for i, vv := range node.values {
			if v < vv {
				node = node.leafs[i]
				f = true
			}
		}
		if !f {
			node = node.leafs[len(node.values)]
		}
	}

	parent := node.parent
SpliterLoop:
	for {
		var left, right []T
		v, left, right = node.insert(v)
		if left == nil && right == nil {
			break SpliterLoop
		}

		if parent == nil {
			parent = newBtreeNode[T](cap(node.values), nil)
			node.splitLeafs(v, left, right)
			r.root = parent
			continue
		}
		node = parent
		node.splitLeafs(v, left, right)
	}

}
