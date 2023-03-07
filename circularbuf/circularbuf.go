package circularbuf

import "errors"

// Buffer is a circular buffer.
// Allows to go to the next and prev value and seek for value.
type Buffer[T any] struct {
	cur int
	inner []T
}
 
// NewBuffer creates new circular buffer that holds any type.
func NewBuffer[T any] (cap int) *Buffer[T] {
	inner := make([]T, 0, cap)
	return &Buffer[T]{cur: -1, inner: inner}
}

// Pos returns current buffer position.
func (b *Buffer[T]) Pos() int {
	return b.cur
}

// Get returns value at current cursor position.
func (b *Buffer[T]) Get() (T, error){
	if b.cur < 0 {
		var t T
		return t, errors.New("empty buffer")
	}
	return b.inner[b.cur], nil
}

// Add adds value at current position.
func (b *Buffer[T]) Add(value T) {
	if b.cur <= 0 {
		b.inner = append(b.inner, value)
		return
	}
	b.inner = append(b.inner[:b.cur], append([]T{value}, b.inner[b.cur+1:]...)...)
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
	if b.cur == 0 {
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