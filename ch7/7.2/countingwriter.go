package countingwriter

import "io"

// strategy:
// something writes
// some state is held about bytes written from inner writer
type WriteCounter struct {
	count  int64
	writer io.Writer
}

func (c *WriteCounter) Write(p []byte) (int, error) {
	n, err := c.writer.Write(p)
	if err != nil {
		return 0, err
	}
	c.count += int64(n)
	return n, nil
}

func CountingWriter(w io.Writer) (io.Writer, *int64) {
	wc := &WriteCounter{writer: w}
	return wc, &wc.count
}
