package matrix

import (
	"fmt"
	"strconv"
	"strings"
)

type matrix struct {
	dataset [][]int
}

type Matrix interface {
	Rows() [][]int
	Cols() [][]int
	Set(r, c, v int) bool
}

func New(s string) (Matrix, error) {
	r := strings.Split(s, "\n")
	dataset := make([][]int, 0, len(r))
	cNum := 0
	cSet := false
	for _, v := range r {
		vT := strings.Trim(v, " ")
		c := strings.Split(vT, " ")
		cols := make([]int, 0, len(c))
		if !cSet {
			cNum = len(c)
			cSet = true
		} else {
			if cNum != len(c) {
				return nil, fmt.Errorf("not consistent length of a column, expected %v, received %v", cNum, len(c))
			}
		}
		for _, w := range c {
			num, err := strconv.Atoi(w)
			if err != nil {
				return nil, err
			}
			cols = append(cols, num)
		}
		dataset = append(dataset, cols)
	}
	return &matrix{dataset}, nil
}

func (m *matrix) Rows() [][]int {
	cloned := make([][]int, 0, len(m.dataset))
	for _, v := range m.dataset {
		cl := make([]int, len(v))
		copy(cl, v)
		cloned = append(cloned, cl)
	}
	return cloned
}

func (m *matrix) Cols() [][]int {
	if len(m.dataset) == 0 {
		return [][]int{}
	}
	cols := make([][]int, len(m.dataset[0]))
	for i := range cols {
		cols[i] = make([]int, len(m.dataset))
	}
	for i, v := range m.dataset {
		for j, w := range v {
			cols[j][i] = w
		}
	}
	return cols
}

func (m *matrix) Set(r, c, v int) bool {
	if r < 0 || c < 0 {
		return false
	}
	if len(m.dataset) == 0 || len(m.dataset[0]) == 0 {
		return false
	}
	if r >= len(m.dataset) || c >= len(m.dataset[0]) {
		return false
	}

	m.dataset[r][c] = v
	return true
}
