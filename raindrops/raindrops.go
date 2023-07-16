package raindrops

import (
	"strconv"
	"strings"
)

func Convert(number int) string {
	var builder strings.Builder
	for _, d := range []struct {
		s string
		v int
	}{
		{
			v: 3,
			s: "Pling",
		},
		{
			v: 5,
			s: "Plang",
		},
		{
			v: 7,
			s: "Plong",
		},
	} {
		if number%d.v == 0 {
			builder.WriteString(d.s)
		}
	}
	if builder.Len() > 0 {
		return builder.String()
	}
	return strconv.Itoa(number)
}
