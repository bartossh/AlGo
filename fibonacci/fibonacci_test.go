package fibonacci

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBasic(t *testing.T) {
	// given
	cases := []struct {
		n, v int
	}{
		{1, 1},
		{2, 1},
		{3, 2},
		{4, 3},
		{5, 5},
		{6, 8},
		{7, 13},
		{8, 21},
		{9, 34},
	}

	// when

	cf := NewConsecutiveFibonacci()

	for _, c := range cases {
		t.Run(fmt.Sprintf("fibonacci n %v should return %v", c.n, c.v), func(t *testing.T) {
			v := cf.SolveNthValue(c.n)
			// then
			assert.Equal(t, c.v, v)
		})
	}
}

func TestLarge(t *testing.T) {
	// given
	cases := []struct {
		n, v int
	}{
		{30, 832040},
		{40, 102334155},
		{50, 12586269025},
		{55, 139583862445},
		{60, 1548008755920},
		{70, 190392490709135},
		{77, 5527939700884757},
		{80, 23416728348467685},
		{85, 259695496911122585},
	}

	// when

	cf := NewConsecutiveFibonacci()

	for _, c := range cases {
		t.Run(fmt.Sprintf("fibonacci n %v should return %v", c.n, c.v), func(t *testing.T) {
			v := cf.SolveNthValue(c.n)
			// then
			assert.Equal(t, c.v, v)
		})
	}
}

func BenchmarkOneConstruct(b *testing.B) {
	// given
	nth, v := 85, 259695496911122585

	// when

	cf := NewConsecutiveFibonacci()

	for n := 0; n < b.N; n++ {
		nv := cf.SolveNthValue(nth)
		assert.Equal(b, v, nv)
	}
}

func BenchmarkManyConstruct(b *testing.B) {
	// given
	nth, v := 85, 259695496911122585

	// when

	for n := 0; n < b.N; n++ {
		cf := NewConsecutiveFibonacci()
		nv := cf.SolveNthValue(nth)
		assert.Equal(b, v, nv)
	}
}
