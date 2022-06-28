package bubblesort

// Sortable constrain types to those that can preserve order based on comparable value
type Sortable interface {
	~int | ~rune | ~int64 | ~byte | ~uint | ~uint16 | ~uint32 | ~uint64 | ~float32 | ~float64
}

// Sort sorts slice using bubble sort algorithm
func Sort[T Sortable](sl []T) {
	sorted := false
	n := len(sl) - 1
	for !sorted {
		sorted = true
		for i := 0; i < n; i++ {
			if sl[i] > sl[i+1] {
				sl[i], sl[i+1] = sl[i+1], sl[i]
				sorted = false
			}
		}
		n -= 1
	}
}
