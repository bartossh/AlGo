package etl

import "strings"

func Transform(i map[int][]string) map[string]int {
	o := make(map[string]int)
	for k, vs := range i {
		for _, l := range vs {
			lowL := strings.ToLower(l)
			o[lowL] = k
		}
	}
	return o

}
