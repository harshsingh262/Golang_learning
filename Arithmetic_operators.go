package main

import (
	"fmt"
)

func main() {
	var a, b string = "foo", "bar"
	fmt.Println(a + b) // Concatenation of two strings

	var c, d float64 = 3.14, 2.71
	fmt.Println(c - d) // Subtraction of two float64 numbers

	var e, f int = 12, 2
	fmt.Println(e * f) // Multiplication of two integers
	fmt.Println(e / f) // Division of two integers
	fmt.Println(e % f) // Modulus of two integers

	var g int = 1
	g++
	fmt.Println(g) // Incrementing g by 1

	var h int = 5
	h--
	fmt.Println(h) // Decrementing h by 1
}
