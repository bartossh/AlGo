package btree

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInsertAtBTreenodeFullCapacity(t *testing.T) {
	// given

	cases := []struct {
		node  *node[int]
		value int
		index int
		nv    int
		left  []int
		right []int
	}{
		{
			node:  &node[int]{nodes: nil, values: []int{1, 2, 3}},
			value: 0,
			index: 0,
			nv:    1,
			left:  []int{0},
			right: []int{2, 3},
		},
		{
			node:  &node[int]{nodes: nil, values: []int{1, 2, 3, 4}},
			value: 0,
			index: 0,
			nv:    2,
			left:  []int{0, 1},
			right: []int{3, 4},
		},
		{
			node:  &node[int]{nodes: nil, values: []int{1, 3, 4}},
			value: 2,
			index: 1,
			nv:    2,
			left:  []int{1},
			right: []int{3, 4},
		},
		{
			node:  &node[int]{nodes: nil, values: []int{1, 3, 4, 5}},
			value: 2,
			index: 1,
			nv:    3,
			left:  []int{1, 2},
			right: []int{4, 5},
		},
		{
			node:  &node[int]{nodes: nil, values: []int{1, 2, 4}},
			value: 3,
			index: 2,
			nv:    2,
			left:  []int{1},
			right: []int{3, 4},
		},
		{
			node:  &node[int]{nodes: nil, values: []int{1, 2, 4, 5}},
			value: 3,
			index: 2,
			nv:    3,
			left:  []int{1, 2},
			right: []int{4, 5},
		},
		{
			node:  &node[int]{nodes: nil, values: []int{1, 2, 3, 4}},
			value: 5,
			index: 4,
			nv:    3,
			left:  []int{1, 2},
			right: []int{4, 5},
		},
		{
			node:  &node[int]{nodes: nil, values: []int{1, 2, 3, 4, 5, 7, 8}},
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
				removedValue, _, _ := c.node.insertAt(c.index, len(c.node.values), c.value)
				// then
				assert.Equal(t, c.nv, removedValue)
				//	assert.Equal(t, c.left, l.values)
				//	assert.Equal(t, c.right, r.values)
			})
	}
}

func TestInsertAtBTreenodeFreeCapacity(t *testing.T) {
	// given

	cases := []struct {
		node  []int
		value int
		index int
		nv    int
		left  *node[int]
		right *node[int]
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
				values := make([]int, len(c.node))
				copy(values, c.node)
				n := &node[int]{nodes: nil, values: values}
				removedValue, l, r := n.insertAt(c.index, len(values)+1, c.value)
				// then
				assert.Equal(t, c.nv, removedValue)
				assert.Equal(t, c.left, l)
				assert.Equal(t, c.right, r)
			})
	}
}

func BenchmarkBTreenodeInsertAtNotExcededCapacity(b *testing.B) {
	for n := 0; n < b.N; n++ {
		v := make([]int, 9, 10)
		n := &node[int]{nodes: nil, values: v}
		n.insertAt(7, len(n.values)+1, 7)
	}
}

func BenchmarkBTreenodeInsertAtExcededCapacity(b *testing.B) {
	for n := 0; n < b.N; n++ {
		v := make([]int, 10, 10)
		n := &node[int]{nodes: nil, values: v}
		n.insertAt(7, len(n.values), 7)
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
			n := node[int]{nodes: nil, values: v}
			var l, r *node[int]
			for _, v := range c.values {
				_, l, r = n.insert(v, c.cap)
			}
			if l != nil && r != nil {
				assert.Equal(t, c.l, l.values)
				assert.Equal(t, c.r, r.values)
			} else {
				assert.Equal(t, c.result, n.values)
			}
		})

	}
}

func BenchmarkInsertValuesNotExceededCapacity5Inserts(b *testing.B) {
	for n := 0; n < b.N; n++ {
		v := make([]int, 0, 5)
		n := node[int]{nodes: nil, values: v}
		for i := 0; i < 5; i++ {
			n.insert(i, 5)
		}
	}
}

func BenchmarkInsertValuesHalfExceededCapacity5Inserts(b *testing.B) {
	for n := 0; n < b.N; n++ {
		v := make([]int, 0, 2)
		n := node[int]{nodes: nil, values: v}
		for i := 0; i < 5; i++ {
			n.insert(i, 2)
		}
	}
}

func BenchmarkInsertValuesNotExcededCapacity100Inserts(b *testing.B) {
	for n := 0; n < b.N; n++ {
		v := make([]int, 0, 100)
		n := node[int]{nodes: nil, values: v}
		for i := 0; i < 100; i++ {
			n.insert(i, 100)
		}
	}
}

func BenchmarkInsertValuesHalfExcededCapacity100Inserts(b *testing.B) {
	for n := 0; n < b.N; n++ {
		v := make([]int, 0, 50)
		n := node[int]{nodes: nil, values: v}
		for i := 0; i < 100; i++ {
			n.insert(i, 50)
		}
	}
}

func TestTraversalBTree(t *testing.T) {
	sizes := []int{2, 3, 4, 5}
	for _, size := range sizes {
		r := New[int](size)
		for i := 0; i < 50; i++ {
			r.Insert(i)
		}
		r.Traversal()
	}
}

func TestFindInBtree(t *testing.T) {
	ordered := make([]int, 1000)
	reversed := make([]int, 1000)
	even := make([]int, 0, 500)
	odd := make([]int, 0, 500)
	for i := 0; i < 1000; i++ {
		ordered[i] = i
		reversed[i] = 1000 - i - 1
		if i%2 == 0 {
			even = append(even, i)
		} else {
			odd = append(odd, i)
		}
	}

	sizes := []int{2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13}
	cases := []struct {
		name      string
		insert    []int
		find      []int
		toBeFound bool
	}{
		{
			name:      "insert in order look in order, success",
			insert:    ordered,
			find:      ordered,
			toBeFound: true,
		},
		{
			name:      "insert in order look in reversed, success",
			insert:    ordered,
			find:      reversed,
			toBeFound: true,
		},
		{
			name:      "insert in reversed look in order, success",
			insert:    reversed,
			find:      ordered,
			toBeFound: true,
		},
		{
			name:      "insert in reversed look reversed, success",
			insert:    reversed,
			find:      reversed,
			toBeFound: true,
		},
		{
			name:      "insert in odd, look odd, success",
			insert:    odd,
			find:      odd,
			toBeFound: true,
		},
		{
			name:      "insert in even, look even, success",
			insert:    even,
			find:      even,
			toBeFound: true,
		},
		{
			name:      "insert in odd, look even, fail",
			insert:    odd,
			find:      even,
			toBeFound: false,
		},
		{
			name:      "insert even, look odd, success",
			insert:    even,
			find:      odd,
			toBeFound: false,
		},
	}

	for _, c := range cases {
		for _, s := range sizes {
			t.Run(fmt.Sprintf("%s, size %v", c.name, s), func(t *testing.T) {
				r := New[int](s)
				for _, v := range c.insert {
					r.Insert(v)
				}
				for _, v := range c.find {
					f := r.Find(v)
					assert.Equal(t, c.toBeFound, f)
				}
			})
		}
	}
}

func BenchmarkInsertBtreeRoot(b *testing.B) {
	for _, size := range []int{1000, 10000, 100000} {
		for _, cap := range []int{2, 3, 4, 5, 6, 7, 8, 9, 10} {
			r := New[int](cap)
			for v := 0; v < size; v++ {
				r.Insert(v)
			}
			b.Run(fmt.Sprintf("size %v cap %v", size, cap), func(b *testing.B) {
				for n := 0; n < b.N; n++ {
					v := int(rand.Int63n(int64(size)))
					r.Find(v)
				}
			})
		}
	}
}
