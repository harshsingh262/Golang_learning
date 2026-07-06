package main

import (
	"fmt"
)

func main() {
	const name     // Return error because const must be initialized
	name = "Harsh" // Return error because const cannot be reassigned
	fmt.Printf("My name is: %s\n", &name)
}
