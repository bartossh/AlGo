package btree

import (
	"fmt"
	"strings"

	"golang.org/x/exp/constraints"
)

type comperable = constraints.Ordered

type prevNode[T comperable] struct {
	prev *prevNode[T]
	node *node[T]
}

// node is a node in btree data structure
type node[T comperable] struct {
	nodes  []*node[T]
	values []T
}

// BTree is a btree data structure
type BTree[T comperable] struct {
	root     *node[T]
	capacity int
}

func New[T comperable](cap int) *BTree[T] {
	n := new[T](cap)
	return &BTree[T]{n, cap}
}

func new[T comperable](capacity int) *node[T] {
	values := make([]T, 0, capacity)
	return &node[T]{nil, values}
}

// Traversal travers the btree nodes and prints to stdout values adding prefix for tree like structure
func (r BTree[T]) Traversal() {
	r.root.traversal(0)
}

// Insert inserts value to the btree structure
func (r *BTree[T]) Insert(v T) {
	node := r.root

	prev := &prevNode[T]{}

	for {

		if len(node.nodes) == 0 {
			break
		}
		i, _ := node.findIdx(v)
		prev = &prevNode[T]{prev, node}
		node = node.nodes[i]
	}

	var idx int
	v, ln, rn := node.insert(v, r.capacity)
	if ln == nil && rn == nil {
		return
	}

	var ok bool
	for {
		if prev.node == nil {
			prev.node = new[T](r.capacity)
			r.root = prev.node
		}

		idx, ok = prev.node.findIdx(v)
		if !ok {
			return
		}
		v, ln, rn = prev.node.arrange(idx, r.capacity, v, ln, rn)

		if ln == nil && rn == nil {
			return
		}
		prev = prev.prev
	}
}

func (n *node[T]) insert(v T, capacity int) (nv T, ln, rn *node[T]) {
	idx, ok := n.findIdx(v)
	if !ok {
		return
	}
	nv, ln, rn = n.insertAt(idx, capacity, v)
	return nv, ln, rn
}

func (n *node[T]) insertAt(i, capacity int, v T) (nv T, ln, rn *node[T]) {
	n.values = append(n.values[:i], append([]T{v}, n.values[i:]...)...)
	if len(n.values) <= capacity {
		return
	}
	half := (len(n.values) - 1) / 2
	nv = n.values[half]
	ln = new[T](capacity)
	rn = new[T](capacity)
	ln.values = make([]T, len(n.values[:half]), capacity)
	rn.values = make([]T, len(n.values[half+1:]), capacity)
	copy(ln.values, n.values[:half])
	copy(rn.values, n.values[half+1:])
	return
}

func (n *node[T]) arrange(idx, capacity int, v T, ln, rn *node[T]) (nv T, nln, nrn *node[T]) {
	if len(n.values) == 0 {
		n.values = append(n.values, v)
		n.nodes = append(n.nodes, []*node[T]{ln, rn}...)
		return
	}

	n.values = append(n.values[:idx], append([]T{v}, n.values[idx:]...)...)
	n.nodes = append(n.nodes[:idx], append([]*node[T]{ln, rn}, n.nodes[idx+1:]...)...)

	if len(n.values) <= capacity {
		return
	}
	half := (len(n.values) - 1) / 2
	nv = n.values[half]

	nln = new[T](capacity)
	nrn = new[T](capacity)

	nln.values = make([]T, len(n.values[:half]), capacity)
	nrn.values = make([]T, len(n.values[half+1:]), capacity)
	copy(nln.values, n.values[:half])
	copy(nrn.values, n.values[half+1:])

	nln.nodes = make([]*node[T], len(n.nodes[:half+1]), capacity+1)
	nrn.nodes = make([]*node[T], len(n.nodes[half+1:]), capacity+1)
	copy(nln.nodes, n.nodes[:half+1])
	copy(nrn.nodes, n.nodes[half+1:])
	return
}

func (n *node[T]) findIdx(v T) (int, bool) {
	idx := 0
	for i, vv := range n.values {
		if v == vv {
			return 0, false
		}
		if v > vv {
			idx = i + 1
			continue
		}
		break
	}

	return idx, true
}

func (n *node[T]) traversal(spcs int) {
	indent := strings.Repeat("-", spcs)
	hasNodes := len(n.nodes) != 0
	for i := 0; i < len(n.values); i++ {
		if hasNodes {
			n.nodes[i].traversal(spcs + 2)
		}
		fmt.Printf("%s%v\n", indent, n.values[i])
	}
	if hasNodes {
		n.nodes[len(n.nodes)-1].traversal(spcs + 2)
	}
}
