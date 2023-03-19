package armstrong

import (
	"fmt"
	"math"
	"strconv"
)

// IsNumber determinetes if number is an Armstrong number
func IsNumber(num int) bool {
	if num == 0 {
		return true
	}
	str := fmt.Sprintf("%v", num)
	acc := 0
	for _, n := range str {
		if v, err := strconv.Atoi(string(n)); err == nil {
			acc += int(math.Pow(float64(v), float64(len(str))))
		} else {
			return false
		}
	}
	if acc == num {
		return true
	}
	return false
}
