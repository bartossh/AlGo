package shellsort

import "github.com/bartossh/AlGo/constrains"

// Sort uses shell sort algorithm works by swiping the value at a given gap and decreasing the gap to 1
func Sort[T constrains.Sortable](sl []T) {
	gap := len(sl) / 2
	for gap > 0 {
		for i := 0; i < gap; i++ {
			insertion(sl, i, gap)
		}
		gap /= 2
	}
}

func insertion[T constrains.Sortable](sl []T, start, gap int) {
	for i := start + gap; i < len(sl); i += gap {
		cur := sl[i]
		pos := i
		for pos >= gap && sl[pos-gap] > cur {
			sl[pos] = sl[pos-gap]
			pos -= gap
		}
		sl[pos] = cur
	}
}
