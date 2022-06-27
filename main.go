package main

import (
	"fmt"
)

func main() {
	//reverse_words.ReverseWords("Burkay Durdu")
	//result := Solution("abc", "bc") // returns true
	//fmt.Println(result)
	//
	//content := "burkay durdu"
	//test := "aeiou"
	//
	//b := []rune(test)
	//a := []rune(content)
	//
	//fmt.Println(b)
	//fmt.Println(a)

	// XOR
	// 0 0 0
	// 0 1 1
	// 1 0 1
	// 1 1 0

	m := 4 // 1011 - 0100

	n := 1 // 0101 - 0001

	m ^= n // 1110 - 0101
	fmt.Println(m)

	//fmt.Println(FindOdd([]int{123, 345, 567, 123, 345}))
}
