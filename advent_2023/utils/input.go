package utils

import "example.com/advent_2023/types"

// ParseInputToMap reads the given input rows and builds a map from it.
// The key of the map are the Vec2 of the individual input positions and the values as the
// characters of the input represented as int32.
func ParseInputToMap(input []string) map[types.Vec2]int32 {
	m := map[types.Vec2]int32{}
	for y, row := range input {
		for x, c := range row {
			m[types.Vec2{X: x, Y: y}] = c
		}
	}
	return m
}

func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}
