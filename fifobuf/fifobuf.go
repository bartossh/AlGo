package fifobuf

// Buffer is a simplistic fifo buffer.
// Fifo stands for Firs In First Out.
type Buffer[T any] struct {
	inner []T
}

// NewBuffer creates new fifo Buffer.
func NewBuffer[T any](cap int) *Buffer[T] {
	inner := make([]T, 0, cap)
	return &Buffer[T]{inner}
}

// Add adds value to the fifo Buffer.
func (cb *Buffer[T]) Add(value T) {
	cb.inner = append(cb.inner, value)
}

// Get gets value from the fifo Buffer.
func (cb *Buffer[T]) Get() (T, bool) {
	if len(cb.inner) == 0 {
		var t T
		return t, false
	}
	v := cb.inner[0]
	cb.inner = cb.inner[1:]
	return v, true
}
