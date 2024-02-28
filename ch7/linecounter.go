package ch7

import (
	"bufio"
	"bytes"
)

// LineCounter implements io.Writer
type LineCounter int

// Write will increment the line count stored at the address of the variable.
func (lc *LineCounter) Write(b []byte) (int, error) {
	var (
		buf  = bytes.NewBuffer(b)
		scnr = bufio.NewScanner(buf)
	)
	var ct int
	for scnr.Scan() {
		ct++
	}
	*lc = LineCounter(ct)
	return ct, nil
}
