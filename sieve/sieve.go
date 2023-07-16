package sieve

func Sieve(limit int) []int {
	if limit < 2 {
		return nil
	}
	if limit == 2 {
		return []int{2}
	}

	result := make([]int, 0)
	candidates := make(map[int]bool, limit-1)
	for i := 2; i <= limit; i++ {
		if candidates[i] {
			continue
		}
		var z int
		j := 2
		for z <= limit {
			z = i * j
			candidates[z] = true
			j++
		}
		if !candidates[i] {
			result = append(result, i)
		}
	}
	return result
}
