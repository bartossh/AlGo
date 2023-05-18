package circularbuf

import "errors"

// Buffer is a circular buffer.
// Allows to go to the next and prev value and seek for value.
type Buffer[T any] struct {
	cur   int
	inner []T
}

// NewBuffer creates new circular buffer that holds any type.
func NewBuffer[T any](cap int) *Buffer[T] {
	inner := make([]T, 0, cap)
	return &Buffer[T]{cur: -1, inner: inner}
}

// Pos returns current buffer position.
func (b *Buffer[T]) Pos() int {
	return b.cur
}

// Get returns value at current cursor position.
func (b *Buffer[T]) Get() (T, error) {
	if b.cur < 0 {
		var t T
		return t, errors.New("empty buffer")
	}
	return b.inner[b.cur], nil
}

// AddPrev adds value at previous position.
func (b *Buffer[T]) AddPrev(value T) {
	switch {
	case b.cur <= 0:
		b.inner = append(b.inner, value)
		b.cur = 0
	case b.cur == 0:
		b.inner = append([]T{value}, b.inner...)
	default:
		b.inner = append(b.inner[:b.cur], append([]T{value}, b.inner[b.cur:]...)...)
	}
}

// AddNext moves cursor forward and adds a value at this position.
func (b *Buffer[T]) AddNext(value T) {
	b.cur++
	switch {
	case b.cur == len(b.inner), b.cur == 0:
		b.inner = append(b.inner, value)
	default:
		b.inner = append(b.inner[:b.cur], append([]T{value}, b.inner[b.cur+1:]...)...)
	}
}

// RemoveSeek removes value at given cursor position.
// Returns value at removed cursor or error otherwise.
func (b *Buffer[T]) RemoveSeek(cur int) (T, error) {
	if cur < 0 || cur >= len(b.inner) {
		var t T
		return t, errors.New("cursor out of buffer range")
	}
	t := b.inner[cur]
	switch {
	case cur == 0:
		b.inner = b.inner[1:]
	case cur == len(b.inner):
		b.inner = b.inner[:len(b.inner)-1]
	default:
		b.inner = append(b.inner[:cur], b.inner[cur+1:]...)
	}

	if cur < b.cur {
		b.cur--
	}
	if b.cur < 0 {
		b.cur = len(b.inner) - 1
	}
	if b.cur >= len(b.inner) {
		b.cur = 0
	}

	return t, nil
}

// Remove removes value from current cursor position.
// Returns removed value.
func (b *Buffer[T]) Remove() T {
	t := b.inner[b.cur]
	switch {
	case b.cur == 0:
		b.inner = b.inner[1:]
	case b.cur == len(b.inner):
		b.inner = b.inner[:len(b.inner)-1]
	default:
		b.inner = append(b.inner[:b.cur], b.inner[b.cur+1:]...)
	}
	if b.cur == len(b.inner) {
		b.cur = 0
	}

	return t

}

// Next sets cursor to the next position and returns value at this position.
func (b *Buffer[T]) Next() (T, error) {
	if b.cur < 0 {
		var t T
		return t, errors.New("empty buffer")
	}
	b.cur++
	if b.cur == len(b.inner) {
		b.cur = 0
	}
	return b.inner[b.cur], nil
}

// Prev sets cursor to the previous position and returns value at this position.
func (b *Buffer[T]) Prev() (T, error) {
	if b.cur < 0 {
		var t T
		return t, errors.New("empty buffer")
	}
	b.cur--
	if b.cur < 0 {
		b.cur = len(b.inner) - 1
	}
	return b.inner[b.cur], nil
}

// Len returns buffer length.
func (b *Buffer[T]) Len() int {
	return len(b.inner)
}

// Seek sets cursor at specific position.
func (b *Buffer[T]) Seek(cur int) error {
	if cur < 0 || cur >= len(b.inner) {
		return errors.New("cursor out of buffer range")
	}
	b.cur = cur
	return nil
}
