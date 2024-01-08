package day_2

import (
	"fmt"
	"testing"

	"gotest.tools/v3/assert"
)

func TestPreInput(t *testing.T) {
	result, err := Solution1("pre_input.txt")
	assert.NilError(t, err)
	assert.Equal(t, result, 8)
}

func TestInputTask1(t *testing.T) {
	result, err := Solution1("input.txt")
	assert.NilError(t, err)
	fmt.Printf("Day 2 task 1 value equals: [ %v ]\n", result)
}

func TestPreInput2(t *testing.T) {
	result, err := Solution2("pre_input.txt")
	assert.NilError(t, err)
	assert.Equal(t, result, 2286)
}

func TestInputTask2(t *testing.T) {
	result, err := Solution2("input.txt")
	assert.NilError(t, err)
	fmt.Printf("Day 2 task 2 value equals: [ %v ]\n", result)
}
