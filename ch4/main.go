package main

import (
	"crypto"
	"fmt"

	"tgpl/assert"
)

func main() {
	h1 := crypto.SHA256.New()
	h2 := crypto.SHA256.New()
	h1.Write([]byte("hello"))
	h2.Write([]byte("world"))
	sum1, sum2 := h1.Sum(nil), h2.Sum(nil)
	fmt.Printf("sum1: %x\n", sum1)
	fmt.Printf("sum2: %x\n", sum1)
	assert.Equal(231, countDifferentBits(sum1, sum2))

	s := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println("original:", s)
	for i := range s {
		fmt.Printf("rotations: %d > %v\n", i+1, rotate(s, i+1))
	}

}
