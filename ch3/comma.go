package main

import (
	"bytes"
	"strings"
)

// comma accepts integer- and float- strings, with an optional sign (+ or -)
func comma(s string) string {
	if len(s) == 0 {
		return s
	}

	var b bytes.Buffer
	if sign := s[0]; sign == '-' || sign == '+' {
		b.WriteByte(sign)
		s = s[1:] // advance after processing the sign
	}

	decIdx := strings.LastIndex(s, ".")
	if decIdx == -1 {
		// no decimal
		b.WriteString(commaInteger(s))
	} else {
		// yes decimal
		b.WriteString(commaInteger(s[:decIdx]))
		b.WriteString(s[decIdx:])
	}

	return b.String()
}

func commaInteger(s string) string {
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
