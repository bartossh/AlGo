package babystepginatstep

import (
	"errors"
	"math"
)

// Solver is solving discrete logarithm problem:
// a^x = b (mod n) , with respect to gcd(a, n) == 1
// with O(sqrt(n)) time complexity. Uses  Baby-step Giant-step algorithm
// Wikipedia reference: https://en.wikipedia.org/wiki/Baby-step_giant-step
// When a is the primitive root modulo n, the answer is unique.
// Otherwise, it will return the smallest positive solution
func Solver(a, b, n uint64) (uint64, error) {
	if b == 1 {
		return n, nil
	}

	table := make(map[uint64]uint64)
	m := uint64(math.Ceil(math.Sqrt(float64(n))))
	var stp uint64 = 1

	var i uint64

	// baby step
	for i = 0; i < m; i++ {
		table[(stp*b)%n] = i
		stp = (stp * a) % n
	}

	// giant step
	var gstp = stp

	for i = m; i <= n; i += m {
		if v, ok := table[stp]; ok {
			return i - v, nil
		}
		stp = (stp * gstp) % n
	}

	return 0, errors.New("cannot find power of \"a\"")
}
