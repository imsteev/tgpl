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

// this was confusing...
// but i think the idea here is:
// let the underlying reader read however many bytes into p
// if p is larger than how much we can continue to read (Reader.remaining),
// then we make sure to only allow reading by at most Reader.remaining bytes.
// at the end, we still return how much the internal reader read, along with
// any error
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
