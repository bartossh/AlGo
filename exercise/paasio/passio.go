package paasio

import (
	"io"
	"sync"
)

type counterRead struct {
	readCount int
	readBytes int64
	locker    *sync.RWMutex
}

type counterWrite struct {
	writeCount int
	writeBytes int64
	locker     *sync.RWMutex
}

// Writer writes values and counts write operations specifics
type Writer struct {
	w io.Writer
	c counterWrite
}

// NewWriteCounter creates instance allowing access WriteCounter methods
func NewWriteCounter(r io.Writer) WriteCounter {
	var writeBytes int64
	locker := &sync.RWMutex{}
	c := counterWrite{0, writeBytes, locker}
	return &Writer{r, c}
}

// WriteCount returns counted total written bytes and write total operation count
func (w *Writer) WriteCount() (n int64, nops int) {
	w.c.locker.RLock()
	defer w.c.locker.RUnlock()
	return w.c.writeBytes, w.c.writeCount
}

func (w *Writer) Write(p []byte) (n int, err error) {
	n, err = w.w.Write(p)
	if err != nil {
		return
	}
	w.c.locker.Lock()
	defer w.c.locker.Unlock()
	w.c.writeBytes += int64(n)
	w.c.writeCount++
	return
}

// Reader reads values and counts read operation specifics
type Reader struct {
	r io.Reader
	c counterRead
}

// NewReadCounter crate instance allowing access ReadCounter methods
func NewReadCounter(r io.Reader) ReadCounter {
	var readBytes int64
	locker := &sync.RWMutex{}
	c := counterRead{0, readBytes, locker}
	return &Reader{r, c}
}

// ReadCount returns counted total written bytes and write total operation count
func (r *Reader) ReadCount() (n int64, nops int) {
	r.c.locker.RLock()
	defer r.c.locker.RUnlock()
	return r.c.readBytes, r.c.readCount
}

func (r *Reader) Read(p []byte) (n int, err error) {
	n, err = r.r.Read(p)
	if err != nil {
		return
	}
	r.c.locker.Lock()
	defer r.c.locker.Unlock()
	r.c.readBytes += int64(n)
	r.c.readCount++
	return
}

// ReaderWriter reads and writes values and counts read and write operation specifics
type ReaderWriter struct {
	r ReadCounter
	w WriteCounter
}

// NewReadWriteCounter crate instance allowing access ReadCounter and WriteCounter methods
func NewReadWriteCounter(rw io.ReadWriter) ReadWriteCounter {
	r := NewReadCounter(rw)
	w := NewWriteCounter(rw)
	return &ReaderWriter{r, w}
}

// ReadCount returns counted total written bytes and write total operation count
func (rw *ReaderWriter) ReadCount() (n int64, nops int) {
	return rw.r.ReadCount()
}

// WriteCount returns counted total written bytes and write total operation count
func (rw *ReaderWriter) WriteCount() (n int64, nops int) {
	return rw.w.WriteCount()
}

func (rw *ReaderWriter) Read(p []byte) (n int, err error) {
	return rw.r.Read(p)
}

func (rw *ReaderWriter) Write(p []byte) (n int, err error) {
	return rw.w.Write(p)
}
