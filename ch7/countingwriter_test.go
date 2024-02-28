package ch7

import (
	"fmt"
	"testing"
)

func TestCountingWriter(t *testing.T) {
	w, totalBytes := CountingWriter(new(WordCounter))
	w.Write([]byte("one two three"))
	fmt.Printf("total bytes written: %d\n", *totalBytes)
	w.Write([]byte("four five six seven"))
	fmt.Printf("total bytes written: %d\n", *totalBytes)
}
