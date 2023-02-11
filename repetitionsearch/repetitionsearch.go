package repetitionsearch

import "math"

// Search makes linear search through the slice of ints looking for the first repeated number
func Search(sl []int) (int, bool) {
	set := make(map[int]struct{})
	for i := range sl {
		v := sl[i]
		if _, ok := set[v]; ok {
			return i, true
		}
		set[v] = struct{}{}
	}
	return 0, false
}

// SearchFast is the high speed search,
// but is limited with max number value in given slice to be les then math.MaxInt
func SearchFast(sl []int) (int, bool) {
	resp := make([]byte, math.MaxInt16)
	for i, v := range sl {
		j := v / 8
		var mask byte = 1 << (v % 8)
		if resp[j]&mask != 0 {
			return i, true
		}
		resp[j] |= mask
	}

	return -1, false
}
