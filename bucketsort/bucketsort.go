package bucketsort

import (
	"github.com/bartossh/AlGo/constrains"
	"github.com/bartossh/AlGo/insertionsort"
)

// Sort using bucket sort algorithm
func Sort[T constrains.Sortable](sl []T) {
	if sl == nil {
		return
	}

	ln := T(len(sl))

	max := T(0)
	for _, v := range sl {
		if v > max {
			max = v
		}
	}

	buckets := make([][]T, int(ln+1))

	for _, v := range sl {
		bucket := buckets[int(ln*v/max)]
		bucket = append(bucket, v)
		buckets[int(ln*v/max)] = bucket
	}

	for _, bucket := range buckets {
		insertionsort.Sort(bucket)
	}

	nsl := make([]T, 0, len(sl))

	for _, bucket := range buckets {
		for _, v := range bucket {
			nsl = append(nsl, v)
		}
	}
	copy(sl, nsl)
}
