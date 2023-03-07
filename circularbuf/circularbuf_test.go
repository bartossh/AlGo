package circularbuf

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const elements = 100000

func TestMovingNextPrev(t *testing.T) {
	b := NewBuffer[int](1000)
	for i := 0; i < elements; i++ {
		b.AddNext(i)
	}
	for i := 0; i < elements; i++ {
		v, err := b.Next()
		assert.Nil(t, err)
		assert.Equal(t, i, v)
	}
	_, err := b.Next()
	assert.Nil(t, err)
	for i := elements - 1; i >= 0; i-- {
		v, err := b.Prev()
		assert.Nil(t, err)
		assert.Equal(t, i, v)
	}
}

func TestSeek(t *testing.T) {
	b := NewBuffer[int](1000)
	for i := 0; i < elements; i++ {
		b.AddNext(i)
	}
	for i := 0; i < elements; i++ {
		b.Seek(i)
	}
}

func TestRemove(t *testing.T) {
	b := NewBuffer[int](1000)
	for i := 0; i < elements; i++ {
		b.AddNext(i)
	}
	for i := elements - 1; i >= 0; i-- {
		b.Remove()
	}

	assert.Equal(t, 0, b.Len())
}

func TestRemoveSeek(t *testing.T) {
	b := NewBuffer[int](1000)
	for i := 0; i < elements; i++ {
		b.AddNext(i)
	}
	for i := elements - 1; i >= 0; i-- {
		v, err := b.RemoveSeek(i)
		assert.Nil(t, err)
		assert.Equal(t, i, v)
	}

	assert.Equal(t, 0, b.Len())
}

func TestAddPrev(t *testing.T) {
	b := NewBuffer[int](1000)
	for i := 0; i < elements; i++ {
		b.AddPrev(i)
	}
	for i := elements - 1; i >= 0; i-- {
		b.Prev()
		v, err := b.Get()
		assert.Nil(t, err)
		assert.Equal(t, i, v)

	}
}

func TestPos(t *testing.T) {
	b := NewBuffer[int](1000)
	for i := 0; i < elements; i++ {
		b.AddNext(i)
	}
	err := b.Seek(0)
	assert.Nil(t, err)
	for i := 0; i < elements; i++ {
		pos := b.Pos()
		assert.Equal(t, i, pos)
		b.Next()
	}
	err = b.Seek(elements - 1)
	assert.Nil(t, err)
	for i := elements - 1; i >= 0; i-- {
		pos := b.Pos()
		assert.Equal(t, i, pos)
		b.Prev()
	}
}

func BenchmarkAddNext(b *testing.B) {
	for n := 0; n < b.N; n++ {
		buf := NewBuffer[int](1000)
		for i := 0; i < elements; i++ {
			buf.AddNext(i)
		}
	}
}

func BenchmarkAddPrev(b *testing.B) {
	for n := 0; n < b.N; n++ {
		buf := NewBuffer[int](1000)
		for i := 0; i < elements; i++ {
			buf.AddPrev(i)
		}
	}
}

func BenchmarkStepNext(b *testing.B) {
	buf := NewBuffer[int](1000)
	for i := 0; i < elements; i++ {
		buf.AddNext(i)
	}
	for n := 0; n < b.N; n++ {
		for i := 0; i < elements; i++ {
			buf.Next()
		}
	}
}

func BenchmarkStepPrev(b *testing.B) {
	buf := NewBuffer[int](1000)
	for i := 0; i < elements; i++ {
		buf.AddNext(i)
	}
	for n := 0; n < b.N; n++ {
		for i := 0; i < elements; i++ {
			buf.Prev()
		}
	}
}
