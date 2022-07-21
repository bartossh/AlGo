package maximumsubarray

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

const maxBenchArrLength = 100_000

func TestMaximumSubArray(t *testing.T) {
	testCases := []struct {
		array  []int
		result int
	}{
		{array: []int{1, 2, 3, 4, 5, 6, 7, 8, 9}, result: 45},
		{array: []int{1, 0, 5, 8}, result: 14},
		{array: []int{-3, -1, -8, -2}, result: -1},
		{array: []int{-4, 3, -2, 5, -8}, result: 6},
		{array: []int{-8}, result: -8},
		{array: []int{8}, result: 8},
	}

	for _, c := range testCases {
		t.Run(fmt.Sprintf("testing for result %v", c.result), func(t *testing.T) {
			r := Solver(c.array)
			assert.Equal(t, c.result, r)
		})
	}
}

func BenchmarkMaximumSubArray(b *testing.B) {
	b.StopTimer()
	arr := make([]int, maxBenchArrLength)
	for i := range arr {
		arr[i] = int(rand.Int63())
	}
	b.StartTimer()

	for n := 0; n < b.N; n++ {
		Solver(arr)
	}
}
