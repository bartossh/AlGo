package day_22

import (
	"testing"

	"gotest.tools/v3/assert"
)

func TestPreInput(t *testing.T) {
	result, err := Solver1("pre_input.txt")
	assert.NilError(t, err)
	assert.Equal(t, result, 5)
}

//func TestTask1(t *testing.T) {
//	result, err := Solver1("input.txt")
//	assert.NilError(t, err)
//	fmt.Printf("Day 22 task 1 result is [ %v ]", result)
//}
