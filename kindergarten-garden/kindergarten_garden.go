package kindergarten

import (
	"errors"
	"sort"
	"strings"
)

type Garden struct {
	inner map[string][]string
}

func NewGarden(diagram string, children []string) (*Garden, error) {
	set := make(map[string]struct{})
	for _, c := range children {
		set[c] = struct{}{}
	}

	if len(set) != len(children) {
		return nil, errors.New("repreting number of names")
	}

	arr := strings.Split(diagram, "\n")
	if len(arr) != 3 {
		return nil, errors.New("wrong number of separators")
	}
	if len(arr[1]) != len(arr[2]) {
		return nil, errors.New("diagram rows of not equal length")
	}

	if len(arr[1])/2 != len(children) {
		return nil, errors.New("number of plants in diagram not matching number of children")
	}
	if len(arr[1])%2 != 0 {
		return nil, errors.New("odd number of plants in diagram")
	}

	childrenCp := make([]string, len(children))
	copy(childrenCp, children)

	sort.Slice(childrenCp, func(i, j int) bool {
		return childrenCp[i] < childrenCp[j]
	})

	inner := make(map[string][]string)
	for _, a := range []string{arr[1], arr[2]} {
		var idx int
		for i := 0; i < len(arr[1]); i++ {
			plant, err := getPlant(rune(a[i]))
			if err != nil {
				return nil, err
			}
			v := inner[childrenCp[idx]]
			v = append(v, plant)
			inner[childrenCp[idx]] = v

			if i%2 != 0 {
				idx++
			}
		}

	}
	return &Garden{inner}, nil
}

func (g *Garden) Plants(child string) ([]string, bool) {
	v, ok := g.inner[child]
	return v, ok
}

func getPlant(r rune) (string, error) {
	switch r {
	case 'R':
		return "radishes", nil
	case 'C':
		return "clover", nil
	case 'G':
		return "grass", nil
	case 'V':
		return "violets", nil
	default:
		return "", errors.New("unknown plant")
	}
}
