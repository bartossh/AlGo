package fifobuf

import (
	"testing"
)

func TestFifoBufCapLow(t *testing.T) {
	consecutive := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	b := NewBuffer[int](1)
	for _, v := range consecutive {
		b.Add(v)
	}

	for _, v := range consecutive {
		vv, ok := b.Get()
		if !ok || vv != v {
			t.Fatal("should not fail")
		}
	}
}

func TestFifoBufCapHigh(t *testing.T) {
	consecutive := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	b := NewBuffer[int](len(consecutive))
	for _, v := range consecutive {
		b.Add(v)
	}

	for _, v := range consecutive {
		vv, ok := b.Get()
		if !ok || vv != v {
			t.Fatal("should not fail")
		}
	}
}

func TestFifoBufNotEnoughElements(t *testing.T) {
	consecutive := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	b := NewBuffer[int](len(consecutive))
	for _, v := range consecutive {
		b.Add(v)
	}

	for i := 0; i < 100; i++ {
		_, ok := b.Get()
		if i >= 10 && ok {
			t.Fatal("should not provide nay value")
		}
	}
}

func BenchmarkFifoSmallCap(b *testing.B) {
	for n := 0; n < b.N; n++ {
		b := NewBuffer[int](100)
		for i := 100000; i > 0; i-- {
			b.Add(i)
		}
		for {
			_, ok := b.Get()
			if !ok {
				break
			}
		}
	}
}

func BenchmarkFifoLargeCap(b *testing.B) {
	for n := 0; n < b.N; n++ {
		b := NewBuffer[int](100000)
		for i := 100000; i > 0; i-- {
			b.Add(i)
		}
		for {
			_, ok := b.Get()
			if !ok {
				break
			}
		}
	}
}
