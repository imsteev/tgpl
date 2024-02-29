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

	x := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(x)
	fmt.Println(rotate(x, 6))
	fmt.Println(x)

}
