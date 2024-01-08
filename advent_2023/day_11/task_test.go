package day_11

import (
	"fmt"
	"testing"

	"gotest.tools/v3/assert"
)

func TestPreInput(t *testing.T) {
	result, err := Solver1("pre_input.txt")
	assert.NilError(t, err)
	assert.Equal(t, result, 374)
}

func TestTask1(t *testing.T) {
	result, err := Solver1("input.txt")
	assert.NilError(t, err)
	fmt.Printf("Day 11 task 1 result is [ %v ]", result)
}

func TestPreInput2(t *testing.T) {
	result, err := Solver2("pre_input.txt", 2)
	assert.NilError(t, err)
	assert.Equal(t, result, 374)
}

func TestPreInput21(t *testing.T) {
	result, err := Solver2("pre_input.txt", 10)
	assert.NilError(t, err)
	assert.Equal(t, result, 1030)
}

func TestPreInput22(t *testing.T) {
	result, err := Solver2("pre_input.txt", 100)
	assert.NilError(t, err)
	assert.Equal(t, result, 8410)
}

func TestTask2(t *testing.T) {
	result, err := Solver2("input.txt", 1000000)
	assert.NilError(t, err)
	fmt.Printf("Day 11 task 2 result is [ %v ]", result)
}
