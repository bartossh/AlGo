package day_5

import (
	"fmt"
	"testing"

	"gotest.tools/v3/assert"
)

func TestTaskPreInput(t *testing.T) {
	result, err := Solver("pre_input.txt")
	assert.NilError(t, err)
	assert.Equal(t, result, 35)
}

func TestTask1(t *testing.T) {
	result, err := Solver("input.txt")
	assert.NilError(t, err)
	fmt.Printf("Day 5 task 1 result value is [ %v ]", result)
}

func TestTask2PreInput(t *testing.T) {
	result, err := Solver2("pre_input.txt")
	assert.NilError(t, err)
	assert.Equal(t, result, 46)
}

func TestTask2(t *testing.T) {
	result, err := Solver2("input.txt")
	assert.NilError(t, err)
	fmt.Printf("Day 5 task 2 result value is [ %v ]", result)
}
