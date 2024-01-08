package day_21

import (
	"fmt"
	"testing"

	"gotest.tools/v3/assert"
)

func TestPreInput(t *testing.T) {
	result, err := Solver1("pre_input.txt", 6)
	assert.NilError(t, err)
	assert.Equal(t, result, 16)
}

func TestTask1(t *testing.T) {
	result, err := Solver1("input.txt", 64)
	assert.NilError(t, err)
	fmt.Printf("Day 21 task 1 result is [ %v ]", result)
}

func TestPreInput2Task2(t *testing.T) {
	pairs := [][2]int{
		{6, 16},
		{10, 50},
		{50, 1594},
		{100, 6536},
		{500, 167004},
		{1000, 668697},
		{5000, 16733044},
	}
	for _, pair := range pairs {
		t.Run(fmt.Sprintf("steps %v plots %v", pair[0], pair[1]), func(t *testing.T) {
			result, err := Solver1("pre_input.txt", pair[0])
			assert.NilError(t, err)
			assert.Equal(t, result, pair[1])
		})
	}
}

func TestTask2(t *testing.T) {
	result, err := Solver1("input.txt", 26501365)
	assert.NilError(t, err)
	fmt.Printf("Day 21 task 2 result is [ %v ]", result)
}
