package main

import (
	"fmt"
)

func main() {
	// Bitwise AND operator (&)
	var x, y int = 12, 25
	z := x & y     // 12 & 25 = 8
	fmt.Println(z) //

	// Bitwise OR operator (|)
	z = x | y      // 12 | 25 = 29
	fmt.Println(z) // 29

	// Bitwise XOR operator (^)
	z = x ^ y      // 12 ^ 25 = 21
	fmt.Println(z) // 21

	// Bitwise Right Shift operator (>>)
	z = x >> 2     // 12 >> 2 = 3
	fmt.Println(z) // 3

	// Bitwise Left Shift operator (<<)
	z = x << 2     // 12 << 2 = 48
	fmt.Println(z) // 48
}
