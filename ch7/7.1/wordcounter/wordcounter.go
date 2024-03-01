package wordcounter

import (
	"bufio"
	"bytes"
)

// what is buffered I/O?
// it's a technique that lets you control when you actually read/write from disk.
// data is held in-memory until some threshold is met.

type WordCounter int

// Write will increment the word count stored at the address of the variable.
func (wc *WordCounter) Write(b []byte) (int, error) {
	var (
		buf  = bytes.NewBuffer(b)
		scnr = bufio.NewScanner(buf)
	)
	var ct int
	scnr.Split(bufio.ScanWords)
	for scnr.Scan() {
		ct++
	}
	*wc += WordCounter(ct)
	return ct, nil
}
