package main

import (
	"bufio"
	"strings"
)

// 7.1: Implement counters for words and for lines.

// WordCounter implements io.Writer
type WordCounter int

func (wc *WordCounter) Write(b []byte) (int, error) {
	var ct int
	r := strings.NewReader(string(b))
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		ct++
	}
	*wc = WordCounter(ct)
	return ct, nil
}


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
