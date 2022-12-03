package bartoszsort

import (
	"fmt"
	"math/rand"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveEmptyAllocs(t *testing.T) {

	cases := []struct {
		given, expected []int
	}{
		{
			given:    []int{1, 2, 3, 4, 0, 0, 0, 5, 6, 0, 0, 0, 7},
			expected: []int{1, 2, 3, 4, 5, 6, 7},
		},
		{
			given:    []int{1, 2, 3, 4, 0, 0, 0, 5, 6, 0, 7},
			expected: []int{1, 2, 3, 4, 5, 6, 7},
		},
		{
			given:    []int{1, 2, 3, 4, 5, 6, 0, 7},
			expected: []int{1, 2, 3, 4, 5, 6, 7},
		},
		{
			given:    []int{0, 0, 0, 0, 1, 2, 3, 4, 5, 6, 0, 7},
			expected: []int{1, 2, 3, 4, 5, 6, 7},
		},
		{
			given:    []int{0, 0, 0, 0, 1, 2, 3, 4, 5, 6, 0, 7, 0, 0, 0},
			expected: []int{1, 2, 3, 4, 5, 6, 7},
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("test case %v", i), func(t *testing.T) {
			result := removeEmptyAllocs(c.given)
			assert.Equal(t, c.expected, result)
		})
	}

}

func BenchmarkRemoveEmptyAllocs(b *testing.B) {
	b.StopTimer()
	sl := make([]int, 1000000)
	v := 0
	for i := range sl {
		if i+1%5 == 0 {
			v++
			sl[i] = v
		}
	}
	b.StartTimer()

	for n := 0; n < b.N; n++ {
		removeEmptyAllocs(sl)
	}
}

func TestSort(t *testing.T) {

	cases := []struct {
		given, expected []int
	}{
		{
			given:    []int{1, 3, 2, 4, 5, 7, 6, 8, 9},
			expected: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		{
			given:    []int{9, 8, 7, 6, 5, 4, 3, 2, 1},
			expected: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		{
			given:    []int{9, 1, 7, 3, 5},
			expected: []int{1, 3, 5, 7, 9},
		},
		{
			given:    []int{8, 6, 2, 4},
			expected: []int{2, 4, 6, 8},
		},
		{
			given:    []int{20, 19, 18, 17, 16, 15, 14, 13, 21, 22, 23, 24, 25, 12, 11, 10, 9, 7, 8, 6, 4, 5, 2, 3, 1},
			expected: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25},
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("test case %v", i), func(t *testing.T) {
			result := Sort(c.given)
			assert.Equal(t, c.expected, result)
		})
	}

}

func BenchmarkSort(b *testing.B) {
	b.StopTimer()
	sl := make([]int, 1000000)
	v := 0
	for i := range sl {
		if i+1%5 == 0 {
			v++
			sl[i] = v
		}
	}
	b.StartTimer()

	for n := 0; n < b.N; n++ {
		Sort(sl)
	}
}

func BenchmarkSortRandom(b *testing.B) {
	b.StopTimer()
	sl := make([]int, 10000)
	for i := range sl {
		v := rand.Intn(1000000)
		sl[i] = v
	}
	b.StartTimer()

	for n := 0; n < b.N; n++ {
		Sort(sl)
	}
}

func BenchmarkSortConsecutive(b *testing.B) {
	b.StopTimer()
	sl := make([]int, 1000000)
	for i := range sl {
		sl[i] = len(sl) - i

	}
	b.StartTimer()

	for n := 0; n < b.N; n++ {
		Sort(sl)
	}
}

func BenchmarkBuildInGolangSort(b *testing.B) {
	b.StopTimer()
	sl := make([]int, 1000000)
	for i := range sl {
		sl[i] = len(sl) - i

	}
	b.StartTimer()

	for n := 0; n < b.N; n++ {
		sort.Ints(sl)
	}
}

func BenchmarkRandomBuildInGolangSort(b *testing.B) {
	b.StopTimer()
	sl := make([]int, 10000)
	for i := range sl {
		v := rand.Intn(1000000)
		sl[i] = v
	}
	b.StartTimer()

	for n := 0; n < b.N; n++ {
		sort.Ints(sl)
	}
}
