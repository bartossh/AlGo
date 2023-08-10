package isbn

import (
	"strconv"
	"strings"
)

func IsValidISBN(isbn string) bool {
	isbn = strings.ReplaceAll(isbn, "-", "")
	sl := []rune(isbn)
	if len(sl) != 10 {
		return false
	}
	var checkSum int64
	for i, r := range sl {
		v, err := strconv.ParseInt(string(r), 10, 64)
		switch {
		case err != nil && r == 'X' && i == 9:
			checkSum += 10
		case err != nil:
			return false
		default:

			checkSum += v * (10 - int64(i))
		}
	}
	return checkSum%11 == 0
}
