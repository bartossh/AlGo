package quicksort

import (
	"testing"

	"github.com/bartossh/AlGo/constrains"
)

func helperIsSorted[T constrains.Sortable](sl []T) bool {
	for i := range sl[:len(sl)-2] {
		if sl[i] >= sl[i+1] {
			return false
		}
	}
	return true
}

func helperCreateSlice(l int) []int {
	sl := make([]int, 0, l)
	for i := l - 1; i >= 0; i-- {
		sl = append(sl, i)
	}
	return sl
}

func TestDecreasingInt(t *testing.T) {
	testSl := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	Sort(testSl)
	if !helperIsSorted(testSl) {
		t.Errorf("expoected %v to be sorted", testSl)
	}
}

func TestDecreasingInt64(t *testing.T) {
	testSl := []int64{10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	Sort(testSl)
	if !helperIsSorted(testSl) {
		t.Errorf("expoected %v to be sorted", testSl)
	}
}

func TestDecreasingRune(t *testing.T) {
	testSl := []rune{10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	Sort(testSl)
	if !helperIsSorted(testSl) {
		t.Errorf("expoected %v to be sorted", testSl)
	}
}

func TestDecreasingFloat32(t *testing.T) {
	testSl := []float32{10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	Sort(testSl)
	if !helperIsSorted(testSl) {
		t.Errorf("expoected %v to be sorted", testSl)
	}
}

func TestDecreasingFloat64(t *testing.T) {
	testSl := []float64{10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	Sort(testSl)
	if !helperIsSorted(testSl) {
		t.Errorf("expoected %v to be sorted", testSl)
	}
}

func TestDecreasingByte(t *testing.T) {
	testSl := []byte{10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	Sort(testSl)
	if !helperIsSorted(testSl) {
		t.Errorf("expoected %v to be sorted", testSl)
	}
}

func BenchmarkSort(b *testing.B) {
	for n := 0; n < b.N; n++ {
		b.StopTimer()
		testSl := helperCreateSlice(10_000)
		b.StartTimer()
		Sort(testSl)
	}
}

func BenchmarkSort100_000(b *testing.B) {
	for n := 0; n < b.N; n++ {
		b.StopTimer()
		testSl := helperCreateSlice(100_000)
		b.StartTimer()
		Sort(testSl)
	}
}
