package day_20

import (
	"fmt"
	"testing"

	"gotest.tools/v3/assert"
)

func TestPreInput(t *testing.T) {
	result, err := Solver1("pre_input.txt")
	assert.NilError(t, err)
	assert.Equal(t, result, 32000000)
}

func TestPreInput2(t *testing.T) {
	result, err := Solver1("pre_input2.txt")
	assert.NilError(t, err)
	assert.Equal(t, result, 11687500)
}

func TestTask1(t *testing.T) {
	result, err := Solver1("input.txt")
	assert.NilError(t, err)
	fmt.Printf("Day 20 task 1 result is [ %v ]", result)
}

func Test2PreInput2(t *testing.T) {
	result, err := Solver2("pre_input2.txt", "output")
	assert.NilError(t, err)
	assert.Equal(t, result, 1)
}

//func TestTask2(t *testing.T) {
//	result, err := Solver2("input.txt", "rx")
//	assert.NilError(t, err)
//	fmt.Printf("Day 20 task 2 result is [ %v ]", result)
//}
