package pangram

import (
	"strings"
)

const a rune = 97

func IsPangram(s string) bool {
	repl := strings.NewReplacer(" ", "", ".", "", ",", "", ";", "", ":", "", "'", "", "?", "",
		"!", "", "@", "", "#", "", "$", "", "%", "", "^", "", "&", "", "*", "", "(", "", ")", "",
		"_", "", "-", "", "+", "", "=", "")
	s = repl.Replace(s)
	s = strings.ToLower(s)
	distinct := make(map[rune]struct{})
	for _, l := range s {
		if l >= a {
			distinct[l] = struct{}{}
		}
	}
	return len(distinct) == 26
}
