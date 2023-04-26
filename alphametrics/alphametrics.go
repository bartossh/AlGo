package alphametics

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var numsOrder = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}

// Solve returns a solution to numbersic puzzle
func Solve(p string) (map[string]int, error) {
	r := make(map[string]int)
	s := strings.ReplaceAll(p, " ", "")
	puzzle := strings.Split(s, "==")
	strings.Trim(puzzle[0], " ")
	strings.Trim(puzzle[1], " ")
	left := strings.Split(puzzle[0], "+")
	leftS := make([][]string, 0)
	right := strings.Split(puzzle[1], "")
	lettersMap := make(map[string]bool)
	for _, v := range left {
		slot := strings.Split(v, "")
		leftS = append(leftS, slot)
		for _, v := range slot {
			lettersMap[v] = true
		}
	}
	for _, v := range right {
		lettersMap[v] = true
	}
	if len(lettersMap) > len(numsOrder) {
		return r, errors.New("not enough digits")
	}
	combinations := combinationsWithReplacement(len(numsOrder), len(lettersMap))
	for _, combination := range combinations {
		if len(combination) != len(lettersMap) {
			continue
		}
		mv, err := mapCombinationToResult(combination, lettersMap)
		if err != nil {
			continue
		}
		l, err := sumLeft(leftS, mv)
		if err != nil {
			continue
		}
		r, err := slotToNumber(right, mv)
		if err != nil {
			continue
		}
		if l == r {
			return mv, nil
		}
	}
	return nil, errors.New("not possible to find proper combination")
}

func slotToNumber(slot []string, mv map[string]int) (int, error) {
	var ns string
	for i, s := range slot {
		if i == 0 && mv[s] == 0 {
			return 0, errors.New("leading zero solution is invalid")
		}
		ns += fmt.Sprintf("%v", mv[s])
	}
	return strconv.Atoi(ns)
}

func sumLeft(l [][]string, mv map[string]int) (int, error) {
	s := 0
	for _, v := range l {
		n, err := slotToNumber(v, mv)
		if err != nil {
			return s, err
		}
		s += n
	}
	return s, nil
}

func mapCombinationToResult(combination []int, lettersMap map[string]bool) (map[string]int, error) {
	if len(combination) != len(lettersMap) {
		return nil, errors.New("something went wrong, combination has wrong length")
	}
	index := 0
	result := make(map[string]int)

	for k := range lettersMap {
		result[k] = combination[index]
		index++
	}
	return result, nil
}

func combinationsWithReplacement(n, m int) [][]int {
	result := make([][]int, 0)
	indices := make([]int, m)
	if !hasRepetitions(indices) {
		result = append(result, indices)
	}
	for {
		var i int
		for i = m - 1; i >= 0; i-- {
			if indices[i] != n-1 {
				break
			}
		}
		if i < 0 {
			break
		}

		indicesI := indices[i]
		for k := i; k < m; k++ {
			indices[k] = indicesI + 1
		}
		indicesCopy := make([]int, len(indices))
		for i, v := range indices {
			indicesCopy[i] = v
		}
		if !hasRepetitions(indicesCopy) {
			f := func(newS []int) {
				pe := make([]int, len(newS))
				for i := range newS {
					pe[i] = newS[i]
				}
				result = append(result, pe)
			}
			Perm(indicesCopy, f)
		}
	}
	return result
}

func rotateSlice(s []int) []int {
	newS := make([]int, len(s))
	for i := range s {
		if i < len(s)-1 {
			newS[i] = s[i+1]
		} else {
			newS[len(s)-1] = s[0]
		}
	}
	return newS
}

// Perm calls f with each permutation of a.
func Perm(a []int, f func([]int)) {
	perm(a, f, 0)
}

func perm(a []int, f func([]int), i int) {
	if i > len(a) {
		f(a)
		return
	}
	perm(a, f, i+1)
	for j := i + 1; j < len(a); j++ {
		a[i], a[j] = a[j], a[i]
		perm(a, f, i+1)
		a[i], a[j] = a[j], a[i]
	}
}

func hasRepetitions(s []int) bool {
	m := make(map[int]bool)
	for _, v := range s {
		m[v] = true
	}
	return !(len(m) == len(s))
}
