package day_4

import (
	"fmt"
	"testing"

	"gotest.tools/v3/assert"
)

func TestPreInput1(t *testing.T) {
	result, err := Solver1("pre_input.txt")
	assert.NilError(t, err)
	assert.Equal(t, result, 13)
}

func TestTask1(t *testing.T) {
	result, err := Solver1("input.txt")
	assert.NilError(t, err)
	fmt.Printf("Day 4 task 1 value equals: [ %v ]\n", result)
}

func TestPreInput2(t *testing.T) {
	result, err := Solver2("pre_input.txt")
	assert.NilError(t, err)
	assert.Equal(t, result, 30)
}

func TestTask2(t *testing.T) {
	result, err := Solver2("input.txt")
	assert.NilError(t, err)
	fmt.Printf("Day 4 task 2 value equals: [ %v ]\n", result)
}
