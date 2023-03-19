package poker

import (
	"errors"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

const (
	deckSize = 5
	pair     = 2
)

var figureMap = map[string]int{
	"J": 11,
	"D": 12,
	"K": 13,
	"A": 14,
}

var colors = map[string]int{
	"♤": 4,
	"♡": 3,
	"♢": 2,
	"♧": 1,
}

type (
	hand struct {
		high  int
		score int
		str   string
	}
)

func valueToPoints(s string) (int, error) {
	if i, err := strconv.Atoi(s); err == nil {
		return i, nil
	}
	if i, ok := figureMap[s]; ok {
		return i, nil
	}
	return 0, errors.New("not a valid card figure")
}

func splitDeck(s string) (map[int][]int, error) {
	list := strings.Split(s, " ")
	if len(list) != deckSize {
		return map[int][]int{}, fmt.Errorf("wrong deck size, expected: %v, received: %v", deckSize, len(list))
	}
	h := make(map[int][]int)
	for _, v := range list {
		r := []rune(v)
		first := string(r[:len(r)-1])
		last := r[len(r)-1:]
		n, err := valueToPoints(string(first))
		if err != nil {
			return map[int][]int{}, err
		}
		cr, ok := colors[string(last)]
		if !ok {
			return map[int][]int{}, errors.New("color char doesn't exists")
		}
		if cards, ok := h[n]; ok {
			h[n] = append(cards, cr)
		} else {
			h[n] = []int{cr}
		}
	}
	return h, nil
}

func calculateScore(h map[int][]int, s string) hand {
	hd := hand{str: s}
	colors := make(map[int]int)
	for _, v := range h {
		for _, c := range v {
			if _, ok := colors[c]; ok {
				colors[c]++
			} else {
				colors[c] = 0
			}
		}
	}

	values := make([]int, 0, len(h))
	for k := range h {
		values = append(values, k)
	}
	sort.Ints(values)
	hd.high = values[len(values)-1]

	if len(h) == 2 {
		hd.score = 7
		return hd
	}

	if len(colors) == 1 {
		for i := range values {
			if i > values[len(values)-1] {
				if values[i] != values[i+1]-1 {
					hd.score = 6
					return hd
				}
			}
		}
		if hd.high == 13 {
			hd.score = 10
			return hd
		}
		hd.score = 9
		return hd
	}

	if len(colors) == 4 {
		for _, v := range h {
			if len(v) == 4 {
				hd.score = 8
				return hd
			}
			if len(v) == 2 {
				hd.score = 2
				return hd
			}
		}
	}

	if len(h) == 3 {
		for _, v := range h {
			if len(v) == 3 {
				hd.score = 4
				return hd
			}

		}
		hd.score = 3
		return hd
	}

	for i := range values {
		if i > values[len(values)-1] {
			if values[i] == values[i+1]-1 {
				hd.score = 1
				return hd
			}
		}
	}
	hd.score = 5

	return hd

}

func BestHand(hands []string) ([]string, error) {
	scores := make([]hand, 0)
	var max hand
	for _, s := range hands {
		h, err := splitDeck(s)
		if err != nil {
			return []string{}, err
		}
		hs := calculateScore(h, s)
		if hs.score > max.score {
			scores = []hand{hs}
		} else if hs.score == max.score && hs.high == max.high {
			scores = append(scores, hs)
		}
	}
	results := make([]string, 0, len(scores))
	for _, score := range scores {
		results = append(results, score.str)
	}
	fmt.Printf("winning hands %#v\n", results)
	return results, nil
}
