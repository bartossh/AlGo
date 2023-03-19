package circular

import "errors"

// Buffer is a circular buffer
type Buffer struct {
	circ []byte
	size int
}

// NewBuffer create ne instance of a Buffer
func NewBuffer(size int) *Buffer {
	circ := make([]byte, 0)
	return &Buffer{
		circ: circ,
		size: size,
	}
}

//ReadByte implements io.ByteReader for Buffer
func (b *Buffer) ReadByte() (byte, error) {
	if len(b.circ) > 0 {
		v := b.circ[0]
		b.circ = b.circ[1:]
		return v, nil

	}
	return 0, errors.New("empty buffer")
}

// WriteByte implements io.ByteWriter for buffer
func (b *Buffer) WriteByte(c byte) error {
	if len(b.circ) < b.size {
		b.circ = append(b.circ, c)
		return nil
	}
	return errors.New("buffer full")
}

// Overwrite overwrites oldest element in Buffer
func (b *Buffer) Overwrite(c byte) {
	err := b.WriteByte(c)
	if err != nil {
		b.circ = append(b.circ[1:], c)
	}
}

// Reset resets the buffer
func (b *Buffer) Reset() {
	b.circ = make([]byte, 0)
}
