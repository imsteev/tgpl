package main

import (
	"bytes"
)

// comma accepts a non-negative, decimal integer string
// it will return the string formatted with commas
func comma(s string) string {
	// break off chunks that should have a comma in front
	firstChunkSize := len(s)
	for firstChunkSize > 3 {
		firstChunkSize -= 3
	}

	var b bytes.Buffer
	b.WriteString(s[:firstChunkSize])
	for i := firstChunkSize; i < len(s); i += 3 {
		b.WriteString(",")
		b.WriteString(s[i : i+3])
	}
	return b.String()
}
