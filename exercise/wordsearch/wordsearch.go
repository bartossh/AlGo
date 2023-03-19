package wordsearch

import (
	"fmt"
	"strings"
)

func match(w1, w2 string) bool {
	if len(w1) != len(w2) {
		return false
	}
	for i := 0; i < len(w1); i++ {
		if w1[i] != w2[i] {
			return false
		}
	}
	return true
}

func contains(w1, w2 string) (bool, int) {
	for i := range w1[:len(w1)-len(w2)+1] {
		if match(w1[i:len(w2)+i], w2) {
			return true, i
		}
	}
	return false, 0
}

func revertSliceStr(words []string) []string {
	reverted := make([]string, len(words))
	for _, w := range words {
		word := revertStr(w)
		reverted = append(reverted, word)
	}
	return reverted
}

func revertStr(str string) string {
	arr := strings.Split(str, "")
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
	return strings.Join(arr, "")
}

func revertSlice(sl []string) []string {
	reverted := make([]string, len(sl))
	for i, s := range sl {
		reverted[len(sl)-i-1] = s
	}
	return reverted
}

func rotate(puzzle []string) []string {
	rotated := make([]string, len(puzzle[0]))
	for _, p := range puzzle {
		for j, v := range p {
			rotated[j] += fmt.Sprintf("%s", string(v))
		}
	}
	return rotated
}

// Solve returns map of all founded words with coordinates
func Solve(words, puzzle []string) (result map[string][2][2]int, err error) {
	result = make(map[string][2][2]int)
	// left to right
	for _, w := range words {
		if len(w) == 0 {
			continue
		}
		for j, p := range puzzle {
			if len(w) <= len(p) {
				if ok, i := contains(p, w); ok {
					location := [2][2]int{{i, j}, {i + len(w) - 1, j}}
					result[w] = location
				}
			}
		}
	}
	// right to left
	reverted := revertSliceStr(words)
	for _, w := range reverted {
		if len(w) == 0 {
			continue
		}
		for j, p := range puzzle {
			if len(w) <= len(p) {
				if ok, i := contains(p, w); ok {
					location := [2][2]int{{i + len(w) - 1, j}, {i, j}}
					result[revertStr(w)] = location
				}
			}
		}
	}
	// top to bottom
	rotated := rotate(puzzle)
	for _, w := range words {
		if len(w) == 0 {
			continue
		}
		for j, p := range rotated {
			if len(w) <= len(p) {
				if ok, i := contains(p, w); ok {
					location := [2][2]int{{j, i}, {j, i + len(w) - 1}}
					result[w] = location
				}
			}
		}
	}
	// bottom to top
	for _, w := range reverted {
		if len(w) == 0 {
			continue
		}
		for j, p := range rotated {
			if len(w) <= len(p) {
				if ok, i := contains(p, w); ok {
					location := [2][2]int{{j, i + len(w) - 1}, {j, i}}
					result[revertStr(w)] = location
				}
			}
		}
	}
	// top left to bottom right
	for _, w := range words {
		if len(w) == 0 {
			continue
		}
		for j, p := range puzzle {
			for z := range p {
				cross := ""
				for k := range w {
					if j+len(w) < len(puzzle) && z+len(w) <= len(p) {
						cross += string(puzzle[k+j][k+z])
					}
				}
				if len(w) <= len(cross) {
					if ok, _ := contains(cross, w); ok {
						location := [2][2]int{{z, j}, {z + len(w) - 1, j + len(w) - 1}}
						result[w] = location
					}
				}
			}
		}
	}

	// top right to bottom left
	for _, w := range reverted {
		if len(w) == 0 {
			continue
		}
		for j, p := range puzzle {
			for z := range p {
				cross := ""
				for k := range w {
					if j+len(w) < len(puzzle) && z+len(w) <= len(p) {
						cross += string(puzzle[k+j][k+z])
					}
				}
				if len(w) <= len(cross) {
					if ok, _ := contains(cross, w); ok {
						location := [2][2]int{{z + len(w) - 1, j + len(w) - 1}, {z, j}}
						result[revertStr(w)] = location
					}
				}
			}
		}
	}

	// bottom right to top left
	puzzleRevert := revertSlice(puzzle)
	for _, w := range words {
		if len(w) == 0 {
			continue
		}
		for j, p := range puzzle {
			for z := range p {
				cross := ""
				for k := range w {
					if j+len(w) < len(puzzleRevert) && z+len(w) <= len(p) {
						cross += string(puzzleRevert[k+j][k+z])
					}
				}
				if len(w) <= len(cross) {
					if ok, _ := contains(cross, w); ok {
						location := [2][2]int{{z, len(puzzle) - j - 1}, {z + len(w) - 1, len(puzzle) - j - len(w)}}
						result[w] = location
					}
				}
			}
		}
	}
	// bottom left to top right
	for _, w := range reverted {
		if len() == 0 {
			continue
		}
		for j, p := range puzzleRevert {
			for z := range p {
				cross := ""
				for k := range w {
					if j+len(w) < len(puzzleRevert) && z+len(w) <= len(p) {
						cross += string(puzzleRevert[k+j][k+z])
					}
				}
				if len(w) <= len(cross) {
					if ok, _ := contains(cross, w); ok {
						location := [2][2]int{{z + len(w) - 1, len(puzzleRevert) - j - len(w)}, {z, len(puzzleRevert) - j - 1}}
						result[revertStr(w)] = location
					}
				}
			}
		}
	}
	if len(result) != len(words) {
		err = fmt.Errorf("Not found ")
	}
	return
}w
