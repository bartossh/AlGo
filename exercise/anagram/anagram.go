package anagram

import (
	"strings"
)

func Detect(sub string, candidates []string) []string {
	result := make([]string, 0)
	subLower := strings.ToLower(sub)
	for _, candidate := range candidates {
		candidateLower := strings.ToLower(candidate)
		if themselves(subLower, candidateLower) {
			continue
		}
		letters := getLettresDistribution(subLower)
		if anagram(candidateLower, letters) {
			result = append(result, candidate)
		}
	}
	return result
}

func getLettresDistribution(sub string) map[string]int {
	letters := make(map[string]int)
	for _, l := range sub {
		s := string(l)
		letters[s]++
	}
	return letters
}

func anagram(candidate string, lettres map[string]int) bool {
	for _, l := range candidate {
		s := string(l)
		if v, ok := lettres[s]; ok {
			if v == 1 {
				delete(lettres, s)
			} else {
				lettres[s]--
			}
		} else {
			return false
		}
	}
	return len(lettres) == 0
}

func themselves(a, b string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
