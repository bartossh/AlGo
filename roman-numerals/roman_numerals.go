package romannumerals

import (
	"fmt"
	"strings"
)

type arabicToRoman struct {
	roman  string
	arabic int
}

func ToRomanNumeral(number int) (string, error) {
	var buf strings.Builder
	if number <= 0 || number >= 4000 {
		return "", fmt.Errorf("number %d is not allowed", number)
	}
	steps := []arabicToRoman{
		{"M", 1000},
		{"CM", 900},
		{"D", 500},
		{"CD", 400},
		{"C", 100},
		{"XC", 90},
		{"L", 50},
		{"XL", 40},
		{"X", 10},
		{"IX", 9},
		{"V", 5},
		{"IV", 4},
		{"I", 1},
	}
	for _, item := range steps {
		for number >= item.arabic {
			buf.WriteString(item.roman)
			number -= item.arabic
		}
	}
	return buf.String(), nil
}
