package collatzconjecture

import (
	"fmt"
	"math"
)

func CollatzConjecture(n int) (int, error) {
	if n >= math.MaxInt/2+1 {
		return 0, fmt.Errorf("number too large")
	}

	if n <= 0 {
		return 0, fmt.Errorf("number must be positive")
	}

	if n == 1 {
		return 0, nil
	}

	var steps int

	for n > 1 {
		if n%2 == 0 {
			n /= 2
			steps++
			continue
		}
		n = 3*n + 1

		steps++
	}

	return steps, nil
}
