package btree

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInsertAtBTreeNodeFullCapacity(t *testing.T) {
	// given

	cases := []struct {
		node  *btreeNode[int]
		value int
		index int
		nv    int
		left  []int
		right []int
	}{
		{
			node:  &btreeNode[int]{leafs: nil, values: []int{1, 2, 3}},
			value: 0,
			index: 0,
			nv:    1,
			left:  []int{0},
			right: []int{2, 3},
		},
		{
			node:  &btreeNode[int]{leafs: nil, values: []int{1, 2, 3, 4}},
			value: 0,
			index: 0,
			nv:    2,
			left:  []int{0, 1},
			right: []int{3, 4},
		},
		{
			node:  &btreeNode[int]{leafs: nil, values: []int{1, 3, 4}},
			value: 2,
			index: 1,
			nv:    2,
			left:  []int{1},
			right: []int{3, 4},
		},
		{
			node:  &btreeNode[int]{leafs: nil, values: []int{1, 3, 4, 5}},
			value: 2,
			index: 1,
			nv:    3,
			left:  []int{1, 2},
			right: []int{4, 5},
		},
		{
			node:  &btreeNode[int]{leafs: nil, values: []int{1, 2, 4}},
			value: 3,
			index: 2,
			nv:    2,
			left:  []int{1},
			right: []int{3, 4},
		},
		{
			node:  &btreeNode[int]{leafs: nil, values: []int{1, 2, 4, 5}},
			value: 3,
			index: 2,
			nv:    3,
			left:  []int{1, 2},
			right: []int{4, 5},
		},
		{
			node:  &btreeNode[int]{leafs: nil, values: []int{1, 2, 3, 4}},
			value: 5,
			index: 4,
			nv:    3,
			left:  []int{1, 2},
			right: []int{4, 5},
		},
		{
			node:  &btreeNode[int]{leafs: nil, values: []int{1, 2, 3, 4, 5, 7, 8}},
			value: 6,
			index: 5,
			nv:    4,
			left:  []int{1, 2, 3},
			right: []int{5, 6, 7, 8},
		},
	}

	for i, c := range cases {
		t.Run(
			fmt.Sprintf("test case %v, insert value %v at index %v", i, c.value, c.index),
			func(t *testing.T) {
				// when
				removedValue, l, r := c.node.insertAt(c.index, c.value)
				// then
				assert.Equal(t, c.nv, removedValue)
				assert.Equal(t, c.left, l)
				assert.Equal(t, c.right, r)
			})
	}
}

func TestInsertAtBTreeNodeFreeCapacity(t *testing.T) {
	// given

	cases := []struct {
		node  []int
		value int
		index int
		nv    int
		left  []int
		right []int
	}{
		{
			node:  []int{1, 2, 3},
			value: 0,
			index: 0,
			nv:    0,
			left:  nil,
			right: nil,
		},
		{
			node:  []int{1, 2, 3, 4},
			value: 0,
			index: 0,
			nv:    0,
			left:  nil,
			right: nil,
		},
		{
			node:  []int{1, 3, 4},
			value: 2,
			index: 1,
			nv:    0,
			left:  nil,
			right: nil,
		},
		{
			node:  []int{1, 3, 4, 5},
			value: 2,
			index: 1,
			nv:    0,
			left:  nil,
			right: nil,
		},
		{
			node:  []int{1, 2, 4},
			value: 3,
			index: 2,
			nv:    0,
			left:  nil,
			right: nil,
		},
		{
			node:  []int{1, 2, 4, 5},
			value: 3,
			index: 2,
			nv:    0,
			left:  nil,
			right: nil,
		},
		{
			node:  []int{1, 2, 3, 4},
			value: 5,
			index: 4,
			nv:    0,
			left:  nil,
			right: nil,
		},
		{
			node:  []int{1, 2, 3, 4, 5, 7, 8},
			value: 6,
			index: 5,
			nv:    0,
			left:  nil,
			right: nil,
		},
	}

	for i, c := range cases {
		t.Run(
			fmt.Sprintf("test case %v, insert value %v at index %v", i, c.value, c.index),
			func(t *testing.T) {
				// when
				node := make([]int, len(c.node), 100)
				copy(node, c.node)
				n := &btreeNode[int]{leafs: nil, values: node}
				removedValue, l, r := n.insertAt(c.index, c.value)
				// then
				assert.Equal(t, c.nv, removedValue)
				assert.Equal(t, c.left, l)
				assert.Equal(t, c.right, r)
			})
	}
}

func BenchmarkBTreeNodeInsertAtFreeCapacity(b *testing.B) {
	for n := 0; n < b.N; n++ {
		node := make([]int, 9, 10)
		n := &btreeNode[int]{leafs: nil, values: node}
		n.insertAt(7, 7)
	}
}

func BenchmarkBTreeNodeInsertAtFullCapacity(b *testing.B) {
	for n := 0; n < b.N; n++ {
		node := make([]int, 10, 10)
		n := &btreeNode[int]{leafs: nil, values: node}
		n.insertAt(7, 7)
	}
}

func TestInsertValues(t *testing.T) {
	cases := []struct {
		name   string
		cap    int
		values []int
		result []int
		l      []int
		r      []int
	}{
		{
			name:   "small capacity not exceeded, reversed input",
			cap:    5,
			values: []int{5, 4, 3, 2, 1},
			result: []int{1, 2, 3, 4, 5},
		},
		{
			name:   "small capacity exceeded, reversed input",
			cap:    5,
			values: []int{6, 5, 4, 3, 2, 1},
			result: nil,
			l:      []int{1, 2},
			r:      []int{4, 5, 6},
		},
		{
			name:   "small capacity exceeded, consecutive input",
			cap:    5,
			values: []int{1, 5, 2, 4, 3, 6},
			result: nil,
			l:      []int{1, 2},
			r:      []int{4, 5, 6},
		},
		{
			name:   "large capacity  not exceeded, shuffled input",
			cap:    10,
			values: []int{1, 5, 2, 4, 3, 6, 9, 0, 8, 7},
			result: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		{
			name:   "large capacity exceeded, shuffled input",
			cap:    9,
			values: []int{1, 5, 2, 4, 3, 6, 9, 0, 8, 7},
			result: nil,
			l:      []int{0, 1, 2, 3},
			r:      []int{5, 6, 7, 8, 9},
		},
	}
	for i, c := range cases {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			v := make([]int, 0, c.cap)
			n := btreeNode[int]{leafs: nil, values: v}
			var l, r []int
			for _, v := range c.values {
				_, l, r = n.insert(v)
			}
			if l != nil && r != nil {
				assert.Equal(t, c.l, l)
				assert.Equal(t, c.r, r)
			} else {
				assert.Equal(t, c.result, n.values)
			}
		})

	}
}
