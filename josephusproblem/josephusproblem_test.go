package josephusproblem

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve(t *testing.T) {
	problems := []struct {
		soldiers, expected uint
	}{
		{41, 19},
		{1, 1},
		{2, 1},
		{14, 13},
		{7, 7},
		{8, 1},
	}
	for _, p := range problems {
		t.Run(fmt.Sprintf("number of soldiers %v, expected position %v", p.soldiers, p.expected), func(t *testing.T) {
			v := Solve(p.soldiers)
			assert.Equal(t, p.expected, v)
		})

	}
}

func BenchmarkSolve(b *testing.B) {
	var soldiers uint = 41
	for n := 0; n < b.N; n++ {
		Solve(soldiers)
	}
}
