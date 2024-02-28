package main

import "fmt"

func main() {
	// integers
	fmt.Println(comma("1"))
	fmt.Println(comma("12"))
	fmt.Println(comma("123"))
	fmt.Println(comma("1234"))
	fmt.Println(comma("12345"))
	fmt.Println(comma("123456"))
	fmt.Println(comma("1234567"))
	fmt.Println(comma("12345678"))
	fmt.Println(comma("123456789"))
	fmt.Println(comma("1123456789"))

	// floats
	fmt.Println(comma("1.123"))
	fmt.Println(comma("12.123"))
	fmt.Println(comma("123.123"))
	fmt.Println(comma("1234.123"))
	fmt.Println(comma("12345.123"))
	fmt.Println(comma("12345."))
	fmt.Println(comma(".707"))

	// signs
	fmt.Println(comma("-.707"))
	fmt.Println(comma("-1.707"))
	fmt.Println(comma("+10101.707"))

	fmt.Println(anagram("", ""))
	fmt.Println(anagram("1", ""))
	fmt.Println(anagram("1", "1"))
	fmt.Println(anagram("12", "21"))
	fmt.Println(anagram("123", "3211"))
	fmt.Println(anagram("1231", "3211")) // true
}
