package main

import (
	"fmt"
)

func main() {
	// Assignment operator (=)
	//var x int = 10
	//var y int
	//y = x
	//fmt.Println(y) // 10

	// Addition assignment operator (+=)
	var x, y int = 10, 20
	x += y         // x = x + y
	fmt.Println(x) // 30

	// Subtraction assignment operator (-=)
	x -= y         // x = x - y
	fmt.Println(x) // 10

	// Multiplication assignment operator (*=)
	x *= y         // x = x * y
	fmt.Println(x) // 200

	// Division assignment operator (/=)
	x /= y         // x = x / y
	fmt.Println(x) // 10

	// Modulus assignment operator (%=)
	x %= y         // x = x % y
	fmt.Println(x) // 10

}
