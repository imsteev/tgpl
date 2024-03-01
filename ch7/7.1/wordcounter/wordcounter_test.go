package wordcounter

import (
	"fmt"
	"testing"
)

func TestWordCounter(t *testing.T) {
	strs := []string{
		"i want another donut",
		"",
		"word",
	}
	for _, s := range strs {
		var c WordCounter
		c.Write([]byte(s))
		fmt.Printf("wordcount(%q): %d\n", s, c)
	}
}
