package day_14

import (
	"fmt"
	"testing"

	"gotest.tools/v3/assert"
)

func TestPreInput(t *testing.T) {
	result, err := Solver1("pre_input.txt")
	assert.NilError(t, err)
	assert.Equal(t, result, 136)
}

func TestTask1(t *testing.T) {
	result, err := Solver1("input.txt")
	assert.NilError(t, err)
	fmt.Printf("Day 14 task 1 result is [ %v ]", result)
}

func BenchmarkTask1(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Solver1("input.txt")
	}
}

func TestPreInput2(t *testing.T) {
	result, err := Solver2("pre_input.txt", 1000000000)
	assert.NilError(t, err)
	assert.Equal(t, result, 64)
}

func TestTask2(t *testing.T) {
	result, err := Solver2("input.txt", 1000000000)
	assert.NilError(t, err)
	fmt.Printf("Day 14 task 2 result is [ %v ]", result)
}
