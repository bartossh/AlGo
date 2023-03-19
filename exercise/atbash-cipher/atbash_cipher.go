package atbash

import (
	"bytes"
	"strings"
)

const (
	alphabet = "abcdefghijklmnopqrstuvwxyz"
	nums     = "1234567890"
)

func Atbash(s string) string {
	var buff bytes.Buffer
	rpl := strings.NewReplacer(" ", "", ".", "", ",", "", "?", "", "!", "", ":", "", ";", "", "'", "")
	s = rpl.Replace(s)
	s = strings.ToLower(s)
	k := 0
	for _, l := range s {
		if k%5 == 0 && k != 0 {
			buff.WriteString(" ")
		}
		k++
		for _, num := range nums {
			if l == num {
				buff.WriteString(string(l))
				continue
			}
		}
		for i, al := range alphabet {
			if al == l {
				j := (len(alphabet) - 1) - i
				buff.WriteString(string(alphabet[j]))
			}
		}

	}
	return buff.String()
}
