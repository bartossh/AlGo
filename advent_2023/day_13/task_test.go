package day_12

import (
	"fmt"
	"testing"

	"gotest.tools/v3/assert"
)

func TestPreInput(t *testing.T) {
	result, err := Solver1("pre_input.txt")
	assert.NilError(t, err)
	assert.Equal(t, result, 405)
}

func TestInputCustom(t *testing.T) {
	result, err := Solver1("input_custom.txt")
	assert.NilError(t, err)
	assert.Equal(t, result, 100)
}

func TestTask1(t *testing.T) {
	result, err := Solver1("input.txt")
	assert.NilError(t, err)
	fmt.Printf("Day 13 task 1 result is [ %v ]", result)
}

func TestPreInput2(t *testing.T) {
	result, err := Solver2("pre_input.txt")
	assert.NilError(t, err)
	assert.Equal(t, result, 400)
}

func TestTask2(t *testing.T) {
	result, err := Solver2("input.txt")
	assert.NilError(t, err)
	fmt.Printf("Day 13 task 2 result is [ %v ]", result)
}
