package cocktailshakersort

import "github.com/bartossh/AlGo/constrains"

// Sort using cocktail shaker sort algorithm
func Sort[T constrains.Sortable](sl []T) {
	if sl == nil {
		return
	}

outer:
	for {
		swapped := false

		for i := range sl[:len(sl)-1] {
			if sl[i] > sl[i+1] {
				sl[i], sl[i+1] = sl[i+1], sl[i]
				swapped = true
			}
		}
		if !swapped {
			break outer
		}

		swapped = false

		for i := len(sl) - 2; i > 0; i-- {
			if sl[i] > sl[i+1] {
				sl[i], sl[i+1] = sl[i+1], sl[i]
				swapped = true
			}
		}

		if !swapped {
			break outer
		}
	}
}
