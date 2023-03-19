package encode

import (
	"fmt"
	"strconv"
	"unicode"
)

// RunLengthEncode encodes string
func RunLengthEncode(s string) string {
	if len(s) < 2 {
		return s
	}
	encoded := ""
	counter := 1
	var last rune
	for i, r := range s {
		if i == 0 {
			last = r
			continue
		}
		if r == last {
			counter++
		} else {
			if counter > 1 {
				encoded += fmt.Sprintf("%v%s", counter, string(last))
				counter = 1
			} else {
				encoded += string(last)
			}
			last = r
		}
	}
	if counter > 1 {
		encoded += fmt.Sprintf("%v%s", counter, string(last))
	} else {
		encoded += string(last)
	}
	return encoded
}

// RunLengthDecode decodes string
func RunLengthDecode(s string) string {
	decoded := ""
	var digitS string
	for _, r := range s {
		if unicode.IsDigit(r) {
			digitS += string(r)
		} else {
			if len(digitS) > 0 {
				if d, err := strconv.Atoi(digitS); err == nil {
					for i := 0; i < d; i++ {
						decoded += string(r)
					}
				}
			} else {
				decoded += string(r)
			}
			digitS = ""
		}
	}
	return decoded
}
