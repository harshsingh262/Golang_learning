package main

import (
	"fmt"
)

func main() {
	// if statement
	var a string = "Happy"
	if a == "Happy" {
		fmt.Println(a)
	}

	// if-else statement
	var fruit string = "grapes"
	if fruit == "apple" {
		fmt.Println("Fruit is apple")
	} else {
		fmt.Println("Fruit is not apple")
	}
}
