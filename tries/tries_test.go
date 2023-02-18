package tries

import (
	"fmt"
	"testing"
	"unicode/utf8"

	"github.com/stretchr/testify/assert"
)

var frazes = []string{
	"doc",
	"dock",
	"docker",
	"docking",
	"docker-compose",
	"dec",
	"deck",
	"decker",
	"december",
	"Doc",
	"Dock",
	"Docker",
	"Docking",
	"Dockuments",
	"a",
	"aaa",
	"aaaa",
}

func TestAddSuccess(t *testing.T) {
	n := New()
	for _, fraze := range frazes {
		ok := n.Insert(fraze)
		assert.True(t, ok)
	}

	for _, fraze := range frazes {
		ok := n.Find(fraze)
		assert.True(t, ok)
	}
}

func TestAddFailureArleadyExixsts(t *testing.T) {
	n := New()
	for _, fraze := range frazes {
		ok := n.Insert(fraze)
		assert.True(t, ok)
	}

	for _, fraze := range frazes {
		ok := n.Find(fraze)
		assert.True(t, ok)
	}

	for _, fraze := range frazes {
		ok := n.Insert(fraze)
		assert.False(t, ok)
	}
}

func TestFindFailureNotInserted(t *testing.T) {
	n := New()
	for _, fraze := range frazes {
		ok := n.Insert(fraze)
		assert.True(t, ok)
	}

	for _, fraze := range frazes {
		ok := n.Find(fraze)
		assert.True(t, ok)
	}

	for _, fraze := range frazes {
		ok := n.Find(fmt.Sprintf("%s%s", fraze, "z"))
		assert.False(t, ok)
	}

	for _, fraze := range frazes {
		ok := n.Find(fmt.Sprintf("%s%s", "z", fraze))
		assert.False(t, ok)
	}
}

func TestDelete(t *testing.T) {
	n := New()
	for _, fraze := range frazes {
		ok := n.Insert(fraze)
		assert.True(t, ok)
	}

	ok := n.Delete(frazes[0])
	assert.True(t, ok)

	ok = n.Find(frazes[0])
	assert.False(t, ok)
}

func BenchmarkInsert(b *testing.B) {
	chars := "1234567890-_=+{}[]:;<>,.?'"
	largeSet := make([]string, 0, len(frazes)*len(chars))
	for _, fraze := range frazes {
		largeSet = append(largeSet, fraze)
		for i, w := 0, 0; i < len(chars); i += w {
			l, width := utf8.DecodeRuneInString(chars[i:])
			w = width
			largeSet = append(largeSet, fmt.Sprintf("%s%s%s", string(l), fraze, string(l)))
		}
	}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		nn := New()
		for _, fraze := range largeSet {
			ok := nn.Insert(fraze)
			assert.True(b, ok)
		}
	}

}

func BenchmarkFind(b *testing.B) {
	chars := "1234567890-_=+{}[]:;<>,.?'"
	largeSet := make([]string, 0, len(frazes)*len(chars))
	for _, fraze := range frazes {
		largeSet = append(largeSet, fraze)
		for i, w := 0, 0; i < len(chars); i += w {
			l, width := utf8.DecodeRuneInString(chars[i:])
			w = width
			largeSet = append(largeSet, fmt.Sprintf("%s%s%s", string(l), fraze, string(l)))
		}
	}

	b.ResetTimer()

	nn := New()
	for _, fraze := range largeSet {
		ok := nn.Insert(fraze)
		assert.True(b, ok)
	}

	for n := 0; n < b.N; n++ {
		for _, fraze := range largeSet {
			ok := nn.Find(fraze)
			assert.True(b, ok)
		}
	}

}

func BenchmarkDelete(b *testing.B) {
	chars := "1234567890-_=+{}[]:;<>,.?'"
	largeSet := make([]string, 0, len(frazes)*len(chars))
	for _, fraze := range frazes {
		largeSet = append(largeSet, fraze)
		for i, w := 0, 0; i < len(chars); i += w {
			l, width := utf8.DecodeRuneInString(chars[i:])
			w = width
			largeSet = append(largeSet, fmt.Sprintf("%s%s%s", string(l), fraze, string(l)))
		}
	}

	b.ResetTimer()

	nn := New()
	for _, fraze := range largeSet {
		ok := nn.Insert(fraze)
		assert.True(b, ok)
	}

	for n := 0; n < b.N; n++ {
		for _, fraze := range largeSet {
			nn.Delete(fraze)
		}
	}

}
