package minesweeper

import (
	"errors"
)

func (b *Board) Count() error {
	cols := len((*b)[0])
	for _, r := range *b {
		if len(r) != cols {
			return errors.New("board is not rectangular shape")
		}
		f := string(r[0])
		if f != "+" && f != "|" {
			return errors.New("wrong board border")
		}
		for _, v := range r {
			c := string(v)
			if c != "*" && c != "|" && c != "+" && c != "-" && c != " " {
				return errors.New("wrong character")
			}
		}
	}
	for i := range *b {
		for j := range (*b)[i] {
			if string((*b)[i][j]) == "*" {
				applyCounter(b, i, j)
			}
		}
	}
	return nil
}

func applyCounter(b *Board, i, j int) {
	for x := i - 1; x <= i+1; x++ {
		for y := j - 1; y <= j+1; y++ {
			v := string((*b)[x][y])
			if v == " " {
				(*b)[x][y] = '1'
			} else if v != "*" && v != "|" && v != "+" && v != "-" {
				(*b)[x][y]++
			}
		}
	}
}
