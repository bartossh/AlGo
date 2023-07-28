package resistorcolor

import "strings"

var codes = map[string]int{
	"black":  0,
	"brown":  1,
	"red":    2,
	"orange": 3,
	"yellow": 4,
	"green":  5,
	"blue":   6,
	"violet": 7,
	"grey":   8,
	"white":  9,
}

// Colors should return the list of all colors.
func Colors() []string {
	res := make([]string, 0, len(codes))
	for k := range codes {
		res = append(res, k)
	}
	return res
}

// ColorCode returns the resistance value of the given color.
func ColorCode(color string) int {
	if v, ok := codes[strings.ToLower(color)]; ok {
		return v
	}
	return -1
}
