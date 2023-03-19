package isogram

import (
	"log"
	"regexp"
	"strings"
)

func removeNotAlphanumerical(s string) string {
	reg, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		log.Fatal(err)
	}
	return strings.ToLower(reg.ReplaceAllString(s, ""))
}

// IsIsogram checks if string is a isogram
func IsIsogram(s string) bool {
	sc := removeNotAlphanumerical(s)
	set := make(map[rune]bool)
	for _, r := range sc {
		if _, ok := set[r]; ok {
			return false
		}
		set[r] = true
	}
	return true
}
