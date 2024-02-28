package ch7

import "io"

// 7.2: Write a function CountingWriter that will return a writer, as well as a
// pointer to an int64 that holds how many bytes have been written at any given time.

// strategy:
// something writes
// some state is held about bytes written from inner writer
type WriteCounter struct {
	b *int64
	w io.Writer
}

func (c *WriteCounter) Write(p []byte) (int, error) {
	n, err := c.w.Write(p)
	if err != nil {
		return 0, err
	}
	*c.b += int64(n)
	return n, nil
}

func CountingWriter(w io.Writer) (io.Writer, *int64) {
	wc := &WriteCounter{b: new(int64), w: w}
	return wc, wc.b
}
