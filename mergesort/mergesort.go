package mergesort

import (
	"github.com/bartossh/AlGo/constrains"
)

// Sort uses merge sort algorithm
func Sort[T constrains.Sortable](sl []T) {
	if sl == nil {
		return
	}
	if len(sl) <= 1 {
		return
	}
	mid := len(sl) / 2
	Sort(sl[:mid])
	Sort(sl[mid:])
	merge(sl, mid)
}

func merge[T constrains.Sortable](sl []T, mid int) {
	lh, rh := make([]T, len(sl[:mid])), make([]T, len(sl[mid:]))
	copy(lh, sl[:mid])
	copy(rh, sl[mid:])
	l, r := 0, 0

	for i := range sl {
		if r == len(rh) || (l < len(lh) && lh[l] < rh[r]) {
			sl[i] = lh[l]
			l++
			continue
		}
		sl[i] = rh[r]
		r++
	}
}
