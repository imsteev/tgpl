package limitreader

import (
	"io"
)

type Reader struct {
	reader    io.Reader
	remaining int
}

func LimitReader(r io.Reader, n int) io.Reader {
	return &Reader{r, n}
}

func (r *Reader) Read(p []byte) (int, error) {
	if r.remaining <= 0 {
		return 0, io.EOF
	}
	if len(p) > r.remaining {
		p = p[:r.remaining]
	}
	n, err := r.reader.Read(p)
	r.remaining -= n
	return n, err
}
