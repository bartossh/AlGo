package luhn

import (
	"strconv"
	"strings"
)

var replacer = strings.NewReplacer(" ", "")

func Valid(sn string) bool {
	sn = replacer.Replace(sn)
	if len(sn) <= 1 {
		return false
	}
	rev := make([]string, len(sn))
	for i, v := range sn {
		rev[(len(sn)-1)-i] = string(v)
	}
	sum := 0
	for i := range rev {
		n, err := strconv.Atoi(rev[i])
		if err != nil {
			return false
		}
		if (i-1)%2 == 0 {
			n *= 2
			if n > 9 {
				n -= 9
			}
		}
		sum += n
	}
	return sum%10 == 0
}
