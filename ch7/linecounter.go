package ch7

import (
	"bufio"
	"strings"
)

// LineCounter implements io.Writer
type LineCounter int

func (lc *LineCounter) Write(b []byte) (int, error) {
	var ct int
	r := strings.NewReader(string(b))
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		ct++
	}
	*lc = LineCounter(ct)
	return ct, nil
}
