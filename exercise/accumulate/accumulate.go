package accumulate

// Accumulate executes passed function on each element aon the list and maps result to new slice
func Accumulate(slice []string, f func(string) string) []string {
	result := make([]string, 0)
	for _, w := range slice {
		r := f(w)
		result = append(result, r)
	}
	return result
}
