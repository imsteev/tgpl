package ch7

import "io"

// 7.2: Write a function CountingWriter that will return a writer, as well as a
// pointer to an int64 that holds how many bytes have been written at any given time.

// strategy:
// something writes
// some state is held about bytes written from inner writer
type WriteCounter struct {
	bytesWritten int64
	writer       io.Writer
}

func (c *WriteCounter) Write(p []byte) (int, error) {
	n, err := c.writer.Write(p)
	if err != nil {
		return 0, err
	}
	c.bytesWritten += int64(n)
	return n, nil
}

func CountingWriter(w io.Writer) (io.Writer, *int64) {
	wc := &WriteCounter{writer: w}
	return wc, &wc.bytesWritten
}
