package day_10

import (
	"fmt"
	"testing"

	"gotest.tools/v3/assert"
)

func TestPreInput(t *testing.T) {
	result, err := Solver1("pre_input.txt")
	assert.NilError(t, err)
	assert.Equal(t, result, 8)
}

func TestTask1(t *testing.T) {
	result, err := Solver1("input.txt")
	assert.NilError(t, err)
	fmt.Printf("Day 10 task 1 result is [ %v ]", result)
}

func TestPreInput21(t *testing.T) {
	result, err := Solver2("pre_input1.txt")
	assert.NilError(t, err)
	assert.Equal(t, result, 4)
}

func TestPreInput22(t *testing.T) {
	result, err := Solver2("pre_input2.txt")
	assert.NilError(t, err)
	assert.Equal(t, result, 4)
}

func TestPreInput23(t *testing.T) {
	result, err := Solver2("pre_input3.txt")
	assert.NilError(t, err)
	assert.Equal(t, result, 8)
}

func TestPreInput24(t *testing.T) {
	result, err := Solver2("pre_input4.txt")
	assert.NilError(t, err)
	assert.Equal(t, result, 10)
}

func TestTask2(t *testing.T) {
	result, err := Solver2("input.txt")
	assert.NilError(t, err)
	fmt.Printf("Day 10 task 2 result is [ %v ]", result)
}
