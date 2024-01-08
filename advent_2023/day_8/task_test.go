package day_8

import (
	"fmt"
	"testing"

	"gotest.tools/v3/assert"
)

func TestPreTasks(t *testing.T) {
	steps1, err := Solver1("pre_input.txt")
	assert.NilError(t, err)
	assert.Equal(t, steps1, 2)
	steps2, err := Solver1("pre_input2.txt")
	assert.NilError(t, err)
	assert.Equal(t, steps2, 6)
}

func TestTask1(t *testing.T) {
	steps, err := Solver1("input.txt")
	assert.NilError(t, err)
	fmt.Printf("Day 8 task 1 results is [ %v ] steps.\n", steps)
}

func TestPreTasks2(t *testing.T) {
	steps1, err := Solver2("pre_input3.txt")
	assert.NilError(t, err)
	assert.Equal(t, steps1, 6)
}

func TestTask2(t *testing.T) {
	steps, err := Solver2("input.txt")
	assert.NilError(t, err)
	fmt.Printf("Day 8 task 2 results is [ %v ] steps.\n", steps)
}
