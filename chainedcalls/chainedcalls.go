package chainedcalls

type Lock[V any] struct {
	value V
}
type Unlock[V any] struct {
	value V
}

// Unlock unlocks for access previously locked. for access.
func (l Lock[V]) Unlock() Unlock[V] {
	return Unlock[V]{l.value}
}

// Write writes data in to unlocked.
func (l Unlock[V]) Write(v V) Unlock[V] {
	return Unlock[V]{v}
}

// Read reads data from the unlocked.
func (l Unlock[V]) Read(v *V) Unlock[V] {
	*v = l.value
	return Unlock[V]{l.value}
}

// Lock locks for access previously unlocked for access.
func (l Unlock[V]) Lock() Lock[V] {
	return Lock[V]{l.value}
}
