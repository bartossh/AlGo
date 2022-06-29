package insertionsort

import "github.com/bartossh/AlGo/constrains"

// Sort slice using in-place insertion sort algorithm.
func Sort[T constrains.Sortable](sl []T) {
	if sl == nil {
		return
	}

	for i := 1; i < len(sl); i++ {
		cur := sl[i]
		j := i - 1
	inner:
		for sl[j] > cur {
			sl[j+1], sl[j] = sl[j], sl[j+1]
			if j == 0 {
				break inner
			}
			j -= 1
		}
	}
}
