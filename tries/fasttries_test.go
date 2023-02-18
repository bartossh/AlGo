package tries

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var smallLettersCollection = []string{
	"doc",
	"docker",
	"dock",
	"do",
	"done",
	"dockering",
	"dockercomposer",
	"dec",
	"decker",
	"deck",
}

func TestFastNodeInsertSuccess(t *testing.T) {
	n := NewFast()

	for _, s := range smallLettersCollection {
		ok, err := n.Insert(s)
		assert.True(t, ok)
		assert.Nil(t, err)
	}
}

func TestFastNodeInsertFailure(t *testing.T) {
	n := NewFast()

	for _, s := range smallLettersCollection {
		ok, err := n.Insert(s)
		assert.True(t, ok)
		assert.Nil(t, err)
	}

	for _, s := range smallLettersCollection {
		ok, err := n.Insert(strings.ToLower(s))
		assert.False(t, ok)
		assert.Nil(t, err)
	}
}

func TestFastNodeInsertError(t *testing.T) {
	n := NewFast()

	for _, s := range smallLettersCollection {
		ok, err := n.Insert(fmt.Sprintf("%s%s", s, "%"))
		assert.False(t, ok)
		assert.NotNil(t, err)
	}
}

func TestFastNodeFindSuccess(t *testing.T) {
	n := NewFast()

	for _, s := range smallLettersCollection {
		ok, err := n.Insert(s)
		assert.True(t, ok)
		assert.Nil(t, err)
	}

	for _, s := range smallLettersCollection {
		ok := n.Find(s)
		assert.True(t, ok)
	}
}

func TestFastNodeFindFailure(t *testing.T) {
	n := NewFast()

	for _, s := range smallLettersCollection {
		ok, err := n.Insert(s)
		assert.True(t, ok)
		assert.Nil(t, err)
	}

	for _, s := range smallLettersCollection {
		ok := n.Find(fmt.Sprintf("%s%s", s, "L"))
		assert.False(t, ok)
	}
}

func TestFastNodeDeleteSuccess(t *testing.T) {
	n := NewFast()

	for _, s := range smallLettersCollection {
		ok, err := n.Insert(s)
		assert.True(t, ok)
		assert.Nil(t, err)
	}

	for _, s := range smallLettersCollection {
		ok := n.Delete(s)
		assert.True(t, ok)
	}
}

func BenchmarkFastNodeInsert(b *testing.B) {
	benchSlice := make([]string, 0, len(smallLettersCollection)*len(smallLettersCollection))
	for _, s1 := range smallLettersCollection {
		for _, s2 := range smallLettersCollection {
			benchSlice = append(benchSlice, fmt.Sprintf("%s%s", s1, s2))
		}
	}

	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		nn := NewFast()
		for _, s := range benchSlice {
			ok, err := nn.Insert(s)
			assert.True(b, ok)
			assert.Nil(b, err)
		}
	}
}

func BenchmarkFastNodeFind(b *testing.B) {
	benchSlice := make([]string, 0, len(smallLettersCollection)*len(smallLettersCollection))
	for _, s1 := range smallLettersCollection {
		for _, s2 := range smallLettersCollection {
			benchSlice = append(benchSlice, fmt.Sprintf("%s%s", s1, s2))
		}
	}

	b.ResetTimer()

	nn := NewFast()
	for _, s := range benchSlice {
		ok, err := nn.Insert(s)
		assert.True(b, ok)
		assert.Nil(b, err)
	}
	for n := 0; n < b.N; n++ {
		for _, s := range benchSlice {
			ok := nn.Find(s)
			assert.True(b, ok)
		}
	}
}

func BenchmarkFastNodeDelete(b *testing.B) {
	benchSlice := make([]string, 0, len(smallLettersCollection)*len(smallLettersCollection))
	for _, s1 := range smallLettersCollection {
		for _, s2 := range smallLettersCollection {
			benchSlice = append(benchSlice, fmt.Sprintf("%s%s", s1, s2))
		}
	}

	b.ResetTimer()

	nn := NewFast()
	for _, s := range benchSlice {
		ok, err := nn.Insert(s)
		assert.True(b, ok)
		assert.Nil(b, err)
	}
	for n := 0; n < b.N; n++ {
		for _, s := range benchSlice {
			nn.Delete(s)
		}
	}
}
