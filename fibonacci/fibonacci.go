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
	v := cf.SolveNthValue(n-1) + cf.SolveNthValue(n-2)
	cf.store[n] = v
	return v
}
