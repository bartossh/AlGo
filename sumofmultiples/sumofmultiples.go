package summultiples

// SumMultiples returns sum of multiples up to the limit
func SumMultiples(limit int, divisors ...int) int {
	multiples := make(map[int]bool, 0)
	for _, d := range divisors {
		if d == 0 {
			continue
		}
		for i := 0; i < limit; i += d {
			if v := multiples[i]; !v {
				multiples[i] = true
			}
		}
	}
	s := 0
	for m := range multiples {
		s += m
	}
	return s
}
