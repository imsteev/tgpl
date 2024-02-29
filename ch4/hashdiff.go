package main

import (
	"crypto"
	"tgpl/assert"

	_ "crypto/sha256"
)

func countDifferentBits(b1, b2 []byte) int {
	var c int
	for i := 0; i < len(b1); i++ {
		diff := b1[i] ^ b2[i]
		c += countBits(diff)
	}
	return c
}

func countBits(n byte) int {
	var c int
	for n > 0 {
		n /= 2
		c++
	}
	return c
}

func main() {
	h1 := crypto.SHA256.New()
	h2 := crypto.SHA256.New()
	h1.Write([]byte("hello"))
	h2.Write([]byte("world"))
	assert.Equal(2, countDifferentBits(h1.Sum(nil), h2.Sum(nil)))
}
