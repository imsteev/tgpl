package countingwriter

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"testing"
	"unicode"
)

type capitalCounter int

// Write will count and store the number of capitalized letters in p
func (c *capitalCounter) Write(p []byte) (int, error) {
	var ct int
	b := bytes.NewReader(p)
	for {
		r, _, err := b.ReadRune() // advances the reader
		if errors.Is(err, io.EOF) {
			break
		}
		if err != nil {
			return 0, err
		}
		if unicode.IsUpper(r) {
			ct++
		}
	}
	*c += capitalCounter(ct)
	return len(p), nil
}

func TestCountingWriter(t *testing.T) {
	capCtr := new(capitalCounter)
	w, n := CountingWriter(capCtr)
	w.Write([]byte("oNe Two ThreE"))
	fmt.Printf("wrote %d bytes\n", *n)
	fmt.Printf("# of capitalized: %d\n", *capCtr)
}
