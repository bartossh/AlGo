package fibonacci

// ConsecutiveFibonacci holds memory about consecutive fibonacci numbers
type ConsecutiveFibonacci struct {
	store map[int]int
}

// NewConsecutiveFibonacci constructs new instance of ConsecutiveFibonacci
func NewConsecutiveFibonacci() ConsecutiveFibonacci {
	store := make(map[int]int)
	return ConsecutiveFibonacci{store: store}
}

// SolveNthValue calculates n'th value of fibonacci sequence
func (cf *ConsecutiveFibonacci) SolveNthValue(n int) int {
	if n <= 2 {
		return 1
	}
	if v, ok := cf.store[n]; ok {
		return v
	}
	v := 1
	before := 1
	for i := 2; i < n; i++ {
		v = v + before
		before = v - before
	}
	cf.store[n] = v
	return v
}
