package secret

import (
	"fmt"
	"strings"
)

var options = []string{"wink", "double blink", "close your eyes", "jump"}

// Handshake returns secret handshake
func Handshake(n uint) []string {
	s := strings.Split(fmt.Sprintf("%b", n), "")
	l := len(s)
	reverse := false
	result := make([]string, 0)
	for i, v := range s {
		ii := l - i
		if v == "1" {
			switch ii {
			case 5:
				reverse = true
			default:
				result = append(result, options[ii-1])
			}
		}
	}
	if !reverse {
		for x, y := 0, len(result)-1; x < y; x, y = x+1, y-1 {
			result[x], result[y] = result[y], result[x]
		}
	}
	return result
}
