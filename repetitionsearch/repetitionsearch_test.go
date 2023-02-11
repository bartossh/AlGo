package repetitionsearch

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNoRepetitionWithSet(t *testing.T) {
	cases := []struct {
		sl []int
	}{
		{
			sl: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		{
			sl: []int{1, 3, 2, 4, 5, 7, 6, 9, 8},
		},
		{
			sl: []int{9, 8, 7, 6, 5, 4, 3, 2, 1},
		},
		{
			sl: []int{9, 8, 7, 6, 5, 4, 3, 2, 1},
		},
		{
			sl: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			i, _ := Search(c.sl)
			assert.Less(t, i, 1)
		})
	}
}

func TestExistRepetitionNotConsecutiveWithSet(t *testing.T) {
	cases := []struct {
		sl     []int
		repIdx int
	}{
		{
			sl:     []int{5, 3, 2, 8, 10, 45, 12, 5, 66},
			repIdx: 7,
		},
		{
			sl:     []int{100, 1, 1222, 1, 22},
			repIdx: 3,
		},
		{
			sl:     []int{1, 7, 9, 12, 34, 2, 4, 5, 9},
			repIdx: 8,
		},
		{
			sl:     []int{123, 45, 78, 34, 28, 48, 49, 45, 6, 2, 3},
			repIdx: 7,
		},
		{
			sl:     []int{0, 0},
			repIdx: 1,
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			i, _ := Search(c.sl)
			assert.Equal(t, c.repIdx, i)
		})
	}
}

func BenchmarkRepetitionSearchWithSet(b *testing.B) {
	b.ReportAllocs()

	b.StopTimer()
	sl := make([]int, 10000)
	for i := range sl {
		sl[i] = i + 1
	}
	b.StartTimer()
	for n := 0; n < b.N; n++ {
		Search(sl)
	}
}

func TestNoRepetitionFast(t *testing.T) {
	cases := []struct {
		sl []int
	}{
		{
			sl: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		{
			sl: []int{1, 3, 2, 4, 5, 7, 6, 9, 8},
		},
		{
			sl: []int{9, 8, 7, 6, 5, 4, 3, 2, 1},
		},
		{
			sl: []int{9, 8, 7, 6, 5, 4, 3, 2, 1},
		},
		{
			sl: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			_, ok := SearchFast(c.sl)
			assert.False(t, ok)
		})
	}
}

func TestExistRepetitionNotConsecutiveFast(t *testing.T) {
	cases := []struct {
		sl     []int
		repIdx int
	}{
		{
			sl:     []int{1, 2, 3, 4, 5, 6, 7, 3, 8, 9},
			repIdx: 7,
		},
		{
			sl:     []int{9, 8, 7, 6, 5, 3, 4, 3, 2, 1},
			repIdx: 7,
		},
		{
			sl:     []int{4, 9, 8, 7, 6, 5, 4, 3, 2, 1},
			repIdx: 6,
		},
		{
			sl:     []int{9, 8, 7, 6, 1, 5, 4, 3, 2, 1},
			repIdx: 9,
		},
		{
			sl:     []int{1, 1},
			repIdx: 1,
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			i, _ := SearchFast(c.sl)
			assert.Equal(t, c.repIdx, i)
		})
	}
}

func BenchmarkRepetitionSearchFast(b *testing.B) {
	b.ReportAllocs()

	b.StopTimer()
	sl := make([]int, 10000)
	for i := range sl {
		sl[i] = i + 1
	}
	b.StartTimer()
	for n := 0; n < b.N; n++ {
		SearchFast(sl)
	}
}
