package ch7

import (
	"fmt"
	"testing"
)

func TestLineCounter(t *testing.T) {
	strs := []string{
		"\n\n\n\n\n",
		"",
		"i like\nto program\n",
	}
	for _, s := range strs {
		var c LineCounter
		c.Write([]byte(s))
		fmt.Printf("linecount(%q): %d\n", s, c)
	}
}
