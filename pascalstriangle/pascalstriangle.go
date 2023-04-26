package pascal

func Triangle(n int) [][]int {
	res := [][]int{{1}}
	if n == 1 {
		return res
	}
	for i := 0; i < n-1; i++ {
		next := nextTriangle(res[i])
		res = append(res, next)
	}
	return res
}

func nextTriangle(ns []int) []int {
	r := make([]int, len(ns)+1, len(ns)+1)
	r[0] = ns[0]
	for i, _ := range r {
		if i > 0 && i < len(ns) {
			r[i] = ns[i-1] + ns[i]
		}
	}
	r[len(ns)] = ns[len(ns)-1]
	return r
}
