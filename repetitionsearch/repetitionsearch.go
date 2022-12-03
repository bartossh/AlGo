package repetitionsearch

// SearchWithSet makes linear search through the slice of ints looking for the first repeated number
func SearchWithSet(sl []int) int {
	set := make(map[int]struct{}, len(sl))
	for i := range sl {
		v := sl[i]
		if _, ok := set[v]; ok {
			return i
		}
		set[v] = struct{}{}
	}
	return -1
}

// SearchWithMutation searches the slice by mutating its values
func SearchWithMutation(slO []int) int {
	slC := make([]int, len(slO))
	copy(slC, slO)
	for i := 0; i < len(slC); i++ {
		idx := abs(slC[i]) - 1
		if slC[idx] < 0 {
			return i
		}
		slC[idx] = slC[idx] * -1
	}
	return -1
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}
