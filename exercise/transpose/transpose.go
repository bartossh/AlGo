package transpose

import "strings"

// Transpose makes matrix rotation
func Transpose(m []string) []string {
	if len(m) == 0 {
		return m
	}
	maxL := 0
	for _, r := range m {
		if len(r) > maxL {
			maxL = len(r)
		}
	}
	res := make([]string, maxL)
	for i := 0; i < maxL; i++ {
		for j := 0; j < len(m); j++ {
			if len(m[j]) > i {
				res[i] += string(m[j][i])
			} else if j+1 < len(m) && m[j+1] != "" {
				res[i] += " "
			}
		}
	}
	if len(m) > 1 {
		for i := len(res) - 1; i > 0; i-- {
			if i == len(res)-1 {
				res[i] = strings.TrimRight(res[i], " ")
			} else if len(res[i+1]) < len(res[i]) {
				res[i] = strings.TrimRight(res[i], " ")

			}
		}
	}
	return res
}
