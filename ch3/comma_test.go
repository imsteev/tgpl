package ch3

import (
	"fmt"
	"testing"
)

func TestComma(t *testing.T) {
	fmt.Println(comma("1234"))
	fmt.Println(comma("12345"))
	fmt.Println(comma("123456"))
	fmt.Println(comma("1234567"))
	fmt.Println(comma("12345678"))
	fmt.Println(comma("123456789"))
	fmt.Println(comma("1123456789"))
}
