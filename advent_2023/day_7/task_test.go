package day_7

import (
	"fmt"
	"testing"

	"gotest.tools/v3/assert"
)

func TestPreinput1(t *testing.T) {
	result, err := Solver1("pre_input.txt")
	assert.NilError(t, err)
	assert.Equal(t, result, 6440)
}

func TestTask1(t *testing.T) {
	result, err := Solver1("input.txt")
	assert.NilError(t, err)
	fmt.Printf("Day 7 task 1 result is [ %v ]\n", result)
}

func TestPreinput2(t *testing.T) {
	result, err := Solver2("pre_input.txt")
	assert.NilError(t, err)
	assert.Equal(t, result, 5905)
}

func TestTask2(t *testing.T) {
	result, err := Solver2("input.txt")
	assert.NilError(t, err)
	fmt.Printf("Day 7 task 2 result is [ %v ]\n", result)
}
