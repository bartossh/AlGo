package diffsquares

import "math"

func SquareOfSum(n int) int {
	res := 0
	for i := 1; i <= n; i++ {
		res += i
	}
	return int(math.Pow(float64(res), 2))
}

func SumOfSquares(n int) int {
	res := 0.0
	for i := 1; i <= n; i++ {
		res += math.Pow(float64(i), 2)
	}
	return int(res)
}

func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
