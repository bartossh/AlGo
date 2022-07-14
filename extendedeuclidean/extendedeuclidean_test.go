package extendedeuclidean

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSimple(t *testing.T) {
	cases := []struct {
		nums     []int
		solution []int
	}{
		{nums: []int{101, 13}, solution: []int{1, 4, -31}},
		{nums: []int{123, 19}, solution: []int{1, -2, 13}},
		{nums: []int{25, 36}, solution: []int{1, 13, -9}},
		{nums: []int{69, 54}, solution: []int{3, -7, 9}},
		{nums: []int{55, 79}, solution: []int{1, 23, -16}},
		{nums: []int{33, 44}, solution: []int{11, -1, 1}},
		{nums: []int{50, 70}, solution: []int{10, 3, -2}},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("case number %v", i), func(t *testing.T) {
			a, b, k := Solve(c.nums[0], c.nums[1])
			assert.Equal(t, a, c.solution[0])
			assert.Equal(t, b, c.solution[1])
			assert.Equal(t, k, c.solution[2])
		})
	}

}
