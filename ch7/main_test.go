package main

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
