package babystepginatstep

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolver(t *testing.T) {
	cases := []struct {
		a, b, n, x uint64
	}{
		{5, 3, 11, 2},
		{3, 83, 100, 9},
		{3, 311401496, 998244353, 178105253},
		{5, 324637211, 1000000007, 976653449},
		{174857, 48604, 150991, 177},
		{912103, 53821, 75401, 2644},
		{448447, 365819, 671851, 23242},
		{220757103, 92430653, 434948279, 862704},
		{176908456, 23538399, 142357679, 14215560},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("Test num %v for, result: %v", i, c.x), func(t *testing.T) {
			result, err := Solver(c.a, c.b, c.n)
			assert.Nil(t, err)
			assert.Equal(t, c.x, result)
		})
	}
}

func BenchmarkSolver(b *testing.B) {
	c := struct {
		a, b, n, x uint64
	}{5, 324637211, 1000000007, 976653449}

	for n := 0; n < b.N; n++ {
		r, err := Solver(c.a, c.b, c.n)
		assert.Nil(b, err)
		assert.Equal(b, c.x, r)
	}
}
