package day_18

import (
	"fmt"
	"testing"

	"gotest.tools/v3/assert"
)

func TestPreInput(t *testing.T) {
	result, err := Solver1("pre_input.txt")
	assert.NilError(t, err)
	assert.Equal(t, result, 62)
}

func TestTask1(t *testing.T) {
	result, err := Solver1("input.txt")
	assert.NilError(t, err)
	fmt.Printf("Day 18 task 1 result is [ %v ]", result)
}

func TestPreInput2(t *testing.T) {
	result, err := Solver2("pre_input.txt")
	assert.NilError(t, err)
	assert.Equal(t, result, 952408144115)
}

func TestTask2(t *testing.T) {
	result, err := Solver2("input.txt")
	assert.NilError(t, err)
	fmt.Printf("Day 18 task 2 result is [ %v ]", result)
}
