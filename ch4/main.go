package main

import (
	"fmt"
)

func main() {
	// h1 := crypto.SHA256.New()
	// h2 := crypto.SHA256.New()
	// h1.Write([]byte("hello"))
	// h2.Write([]byte("world"))
	// assert.Equal(2, countDifferentBits(h1.Sum(nil), h2.Sum(nil)))

	s := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println("original:", s)
	for i := range s {
		fmt.Printf("rotations: %d > %v\n", i+1, rotate(s, i+1))
	}

}
