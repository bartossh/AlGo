package pythagorean

import (
	"fmt"
	"sort"
)

// Triplet holds three numbers that are subject to pythagorean triplet calculation. The three elements of each returned triplet must be in order t[0] <= t[1] <= t[2], and the list of triplets must be in lexicographic
type Triplet [3]int

// Range returns a list of all Pythagorean triplets with sides in the range min to max inclusive
func Range(min, max int) []Triplet {
	triplets := make(map[string]Triplet, 0)
	for i := max; i >= min+2; i-- {
		for j := max - 1; j >= min+1; j-- {
			for k := max - 2; k >= min; k-- {
				t := Triplet{i, j, k}
				t.sort()
				if t.isPythagoreanTriplet() {
					ss := fmt.Sprintf("%v%v%v", t[0], t[1], t[2])
					if _, ok := triplets[ss]; !ok {
						triplets[ss] = t
					}
				}
			}
		}
	}
	result := make([]Triplet, 0)
	for _, v := range triplets {
		result = append(result, v)
	}
	sortTripletSlice(result)
	return result
}

// Sum returns a list of all Pythagorean triplets where the sum a+b+c (the perimeter) is equal to p
func Sum(p int) []Triplet {
	maxN := p/2 - 1
	triplets := make(map[string]Triplet, 0)
	for i := maxN; i > 2; i-- {
		for j := maxN - 1; j > 1; j-- {
			for k := maxN - 2; k > 0; k-- {
				t := Triplet{i, j, k}
				t.sort()
				if t.isPythagoreanTriplet() {
					s := t.sum()
					ss := fmt.Sprintf("%v%v%v", t[0], t[1], t[2])
					if _, ok := triplets[ss]; !ok && s == p {
						triplets[ss] = t
					}
				}
			}
		}
	}
	result := make([]Triplet, 0)
	for _, v := range triplets {
		result = append(result, v)
	}
	sortTripletSlice(result)
	return result
}

func (t *Triplet) isPythagoreanTriplet() bool {
	if (*t)[0]*(*t)[0]+(*t)[1]*(*t)[1] == (*t)[2]*(*t)[2] {
		return true
	}
	return false
}

func (t Triplet) sum() int {
	return t[0] + t[1] + t[2]
}

func (t *Triplet) sort() {
	s := (*t)[:]
	sort.Ints(s)
	t = &Triplet{s[0], s[1], s[2]}
}

func sortTripletSlice(ts []Triplet) {
	sort.Slice(ts, func(i, j int) bool {
		return ts[i][0] < ts[j][0]
	})
}
