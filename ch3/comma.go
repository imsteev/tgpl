package ch3

import (
	"bytes"
)

// comma accepts a non-negative, decimal integer string
// it will return the string formatted with commas

// THIS CAN BE DONE WITH MODULO %

// 1234 ->   [1   , 234]
// 12345 ->  [12  , 345]
// 123456 -> [123 , 456]
func comma(s string) string {
	if len(s) < 3 {
		return s
	}
	var b bytes.Buffer

	// get first group
	r := len(s) % 3
	if r != 0 {
		b.WriteString(s[:r])
	} else {
		b.WriteString(s[:3])
	}
	for i := b.Len(); i < len(s); i += 3 {
		b.WriteString(",")
		b.WriteString(s[i : i+3])
	}
	return b.String()
}
