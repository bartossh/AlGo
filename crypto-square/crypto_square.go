package cryptosquare

import (
	"math"
	"strings"
	"unicode"
)

func Encode(pt string) string {
	s0 := strings.ToLower(pt)
	s1 := strings.ReplaceAll(s0, " ", "")
	letters := []rune(s1)
	norm := make([]rune, 0, len(letters))
	for _, r := range letters {
		if unicode.IsLetter(r) || unicode.IsNumber(r) {
			norm = append(norm, r)
		}
	}

	n := math.Sqrt(float64(len(norm)))

	var c, r int
	switch {
	case math.Trunc(n) == n:
		c, r = int(n), int(n)
	case math.Round(n) > n:
		c, r = int(n)+1, int(n)+1
	default:
		c, r = int(n)+1, int(n)
	}

	diff := c*r - len(norm)
	for i := 0; i < diff; i++ {
		norm = append(norm, ' ')
	}

	var buf strings.Builder
	for ci := 0; ci < c; ci++ {
		for ri := 0; ri <= r; ri++ {
			if ri == r {
				if ci == c-1 {
					continue
				}
				buf.WriteRune(' ')
				continue
			}
			buf.WriteRune(norm[ci+(ri*c)])
		}
	}

	return buf.String()
}
