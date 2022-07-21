package maximumsubarray

// Solver solves maximum sub array sum
func Solver(arr []int) int {
	mem := make([]int, len(arr))
	mem[0] = arr[0]
	result := arr[0]

	for i := 1; i < len(arr); i++ {
		if mem[i-1] > 0 {
			mem[i] = mem[i-1] + arr[i]
		} else {
			mem[i] = arr[i]
		}
		result = max(result, mem[i])
	}
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
