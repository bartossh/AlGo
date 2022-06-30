package quicksort

import "github.com/bartossh/AlGo/constrains"

// Sort uses quick sort algorithm
func Sort[T constrains.Sortable](sl []T) {
	sort(sl, 0, len(sl)-1)
}

func sort[T constrains.Sortable](sl []T, l, h int) {
	if l < h {
		p := partition(sl, l, h)
		sort(sl, l, p-1)
		sort(sl, p+1, h)
	}
}

func partition[T constrains.Sortable](sl []T, l, h int) int {
	pivot := h
	i := l - 1
	j := h

loop:
	for {
		i++
		for sl[i] < sl[pivot] {
			i++
		}
		j--
		for j >= 0 && sl[j] > sl[pivot] {
			j--
		}
		if i >= j {
			break loop
		}
		sl[i], sl[j] = sl[j], sl[i]
	}
	sl[i], sl[pivot] = sl[pivot], sl[i]
	return i
}
