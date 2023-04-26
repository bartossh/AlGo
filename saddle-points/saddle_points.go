package matrix

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type (
	Pair   struct{ x, y int }
	Matrix [][]int
)

func New(str string) (*Matrix, error) {
	if str == "" {
		return nil, errors.New("cannot create matrix from empty string")
	}
	strs := strings.Split(str, "\n")
	mtx := make(Matrix, 0, len(strs))
	rows := 0
	for i, v := range strs {
		nums := strings.Split(v, " ")
		if len(nums) == 0 {
			return nil, fmt.Errorf("row %v is empty", i)
		}
		if rows == 0 {
			rows = len(nums)
		}
		if len(nums) != rows {
			return nil, fmt.Errorf("row %v has different length then other rows", i)
		}
		numi := make([]int, 0, len(nums))
		for _, n := range nums {
			ni, err := strconv.Atoi(n)
			if err != nil {
				return nil, err
			}
			numi = append(numi, ni)
		}

		mtx = append(mtx, numi)
	}
	return &mtx, nil
}

func (m Matrix) Saddle() []Pair {
	res := make([]Pair, 0)

	for i, v := range m {
		for j, vv := range v {
			if biggest(vv, v) {
				if smallest(vv, column(j, m)) {
					res = append(res, Pair{i, j})
				}
			}
		}
	}
	return res
}

func column(i int, l [][]int) []int {
	col := make([]int, 0, len(l))
	for _, v := range l {
		col = append(col, v[i])
	}
	return col
}

func smallest(v int, l []int) bool {
	max := 1 << 62
	smL := make([]int, 0)
	for _, n := range l {
		if n == max {
			smL = append(smL, n)
		} else if n < max {
			max = n
			smL = []int{n}
		}
	}
	for _, s := range smL {
		if s == v {
			return true
		}
	}
	return false
}

func biggest(v int, l []int) bool {
	min := 0
	smL := make([]int, 0)
	for _, n := range l {
		if n == min {
			smL = append(smL, n)
		} else if n > min {
			min = n
			smL = []int{n}
		}
	}
	for _, s := range smL {
		if s == v {
			return true
		}
	}
	return false
}
