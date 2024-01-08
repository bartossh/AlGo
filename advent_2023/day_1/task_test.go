package day_1

import (
	"fmt"
	"testing"

	"gotest.tools/v3/assert"
)

func TestPreInput(t *testing.T) {
	result, err := Solve("pre_input.txt")
	assert.NilError(t, err)
	assert.Equal(t, result, int64(142))
}

func TestInputTask1(t *testing.T) {
	result, err := Solve("input.txt")
	assert.NilError(t, err)
	fmt.Printf("Day 1 task 1 value equals: [ %v ]\n", result)
}

func TestInputPreTask2(t *testing.T) {
	result, err := Solve2("pre_input_2.txt")
	assert.NilError(t, err)
	assert.Equal(t, result, int64(281))
}

func TestMyInput(t *testing.T) {
	result, err := Solve2("my_input.txt")
	assert.NilError(t, err)
	fmt.Printf("My input value equals: [ %v ]\n", result)
}

func TestInputTask2(t *testing.T) {
	result, err := Solve2("input.txt")
	assert.NilError(t, err)
	fmt.Printf("Day 1 task 2 value equals: [ %v ]\n", result)
}
