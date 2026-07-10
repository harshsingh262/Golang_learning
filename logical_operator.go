package main

import (
	"fmt"
)

func main() {
	// Logical AND operator (&&)
	var x, y int = 10, 20
	fmt.Println((x < 100) && (x < 200)) // true
	fmt.Println((x < 300) && (x < 0))   // false

	// Logical OR operator (||)

	fmt.Println((x < 0) || (x < 20)) // true
	fmt.Println((x < 0) || (x > 20)) // false

	// Logical NOT operator (!)

	fmt.Println(!(x > y)) // true
	fmt.Println(!(true))  // false
	fmt.Println(!(false)) // true
}
