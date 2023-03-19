package grains

import (
	"errors"
)

// Square calculates number of grains of given square
func Square(s int) (uint64, error) {
	if s == 0 {
		return 0, errors.New("cannot square value 0")
	}
	if s < 0 {
		return 0, errors.New("no negative square allowed")
	}
	if s > 64 {
		return 0, errors.New("no square bigger than 64 allowed")
	}
	if s == 1 {
		return 1, nil
	}
	var v uint64 = 2
	if s == 2 {
		return v, nil
	}
	for i := 2; i < s; i++ {
		v = v * 2
	}
	return v, nil
}

// Total calculates total number of grains on the chessboard
func Total() uint64 {
	var total uint64 = 0
	for i := 1; i <= 64; i++ {
		t, _ := Square(i)
		total += t
	}
	return total
}
