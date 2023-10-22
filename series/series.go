package series

func All(n int, s string) []string {
	if n > len(s) {
		return []string{}
	}
	result := make([]string, 0, len(s)-n+1)
	for i := 0; i < len(s)-n+1; i++ {
		result = append(result, s[i:i+n])
	}
	return result
}

func UnsafeFirst(n int, s string) string {
	return s[:n]
}
