package rotationalcipher

import (
	"strings"
	"unicode"
)

func RotationalCipher(plain string, shiftKey int) string {
	var builder strings.Builder
	for _, r := range plain {
		switch {
		case unicode.IsLower(r):
			r += rune(shiftKey)
			if r > 122 {
				r = r - 26
			}
			builder.WriteRune(r)
		case unicode.IsUpper(r):
			r += rune(shiftKey)
			if r > 90 {
				r -= 26
			}
			builder.WriteRune(r)
		default:
			builder.WriteRune(r)
		}
	}

	return builder.String()
}
