package brackets

import "fmt"

func isBracket(r rune) bool {
	switch r {
	case '{':
		return true
	case '}':
		return true
	case '[':
		return true
	case ']':
		return true
	case '(':
		return true
	case ')':
		return true
	}
	return false
}

func isEven(n int) bool {
	return n%2 == 0
}

func getOpposite(r rune) rune {
	switch r {
	case '}':
		return '{'
	case ']':
		return '['
	case ')':
		return '('
	}
	return ' '
}

// Bracket checks is string has all bracketmattching
func Bracket(s string) bool {
	bracketMap := make(map[rune][]int)
	var sCleared string
	for _, l := range s {
		if isBracket(l) {
			sCleared += fmt.Sprintf("%v", string(l))
		}
	}
	for i, l := range sCleared {
		opposite := getOpposite(l)
		v, ok := bracketMap[opposite]
		if ok {
			if len(v) > 0 && (v[len(v)-1]+1 == i || isEven(v[len(v)-1]) != isEven(i)) {
				v = v[:len(v)-1]

				bracketMap[opposite] = v
			} else {
				return false
			}
		} else {
			ve, oke := bracketMap[l]
			if oke {
				ve = append(ve, i)
				bracketMap[l] = ve
			} else {
				bracketMap[l] = []int{i}
			}
		}

	}
	for _, v := range bracketMap {
		if len(v) != 0 {
			return false
		}
	}

	return true
}
