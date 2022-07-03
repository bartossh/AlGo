package bubblesort

import "github.com/bartossh/AlGo/constrains"

// Sort sorts slice using bubble sort algorithm
func Sort[T constrains.Sortable](sl []T) {
	if sl == nil {
		return
	}

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
		n--
	}
}
