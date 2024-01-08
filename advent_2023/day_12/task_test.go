package day_12

import (
	"fmt"
	"testing"

	"gotest.tools/v3/assert"
)

func TestPreInputOneLine(t *testing.T) {
	result, err := Solver1("pre_input_custom.txt")
	assert.NilError(t, err)
	assert.Equal(t, result, 3)
}

func TestPreInput(t *testing.T) {
	result, err := Solver1("pre_input.txt")
	assert.NilError(t, err)
	assert.Equal(t, result, 21)
}

func TestTask1(t *testing.T) {
	result, err := Solver1("input.txt")
	assert.NilError(t, err)
	fmt.Printf("Day 12 task 1 result is [ %v ]", result)
}

func TestPreInput2(t *testing.T) {
	result, err := Solver2("pre_input.txt", 5)
	assert.NilError(t, err)
	assert.Equal(t, result, 525152)
}

func TestTask2(t *testing.T) {
	result, err := Solver2("input.txt", 5)
	assert.NilError(t, err)
	fmt.Printf("Day 12 task 2 result is [ %v ]", result)
}
