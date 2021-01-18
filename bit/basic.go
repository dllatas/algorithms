package bit

//package main

import (
	"fmt"
)

func main() {
	var a uint32 = 8
	var b uint32 = 5

	fmt.Printf("a: %b\n", a)
	fmt.Printf("b: %b\n", b)

	// basic operators
	fmt.Printf("a+b: %b\n", a+b)
	fmt.Printf("a-b: %b\n", a-b)
	fmt.Printf("a*b: %b\n", a*b)

	// advanced operators
	// &   -> AND
	fmt.Printf("a&b: %b\n", a&b)
	fmt.Printf("a&b: %b\n", 5&5)
	fmt.Printf("a&b: %b\n", 1&1)
	fmt.Printf("a&b: %b\n", 0&1)
	// |   -> OR
	fmt.Printf("a|b: %b\n", a|b)
	fmt.Printf("a|b: %b\n", 5|5)
	fmt.Printf("a|b: %b\n", 1|1)
	fmt.Printf("a|b: %b\n", 0|1)
	// ^   -> XOR
	// ~   -> NOT
	// &^  -> AND NOT

	// <<  -> left shift
	// >>  -> right shift
}
