package main

import (
	"fmt"
)

func main() {
	var i int = 800
	switch i {
	case 10:
		fmt.Println("i is 10")

	case 100, 200:
		fmt.Println("i is either 100 or 200")

	default:
		fmt.Println("i is neither 0, 100 or 200")

	}

	var x int = 10
	switch x {
	case -5:
		fmt.Println("-5")

	case 10:
		fmt.Println("10")
		fallthrough

	case 20:
		fmt.Println("20")
		fallthrough

	default:
		fmt.Println("default")
	}

}
