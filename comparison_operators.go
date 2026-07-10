package main

import (
	"fmt"
)

func main() {
	var city string = "kolkata"
	var city2 string = "calcutta"
	fmt.Println(city == city2) // false
	fmt.Println(city != city2) // true

	var a, b int = 5, 10
	fmt.Println(a < b) // true

	var c, d int = 10, 10
	fmt.Println(c <= d) // true

	var e, f int = 20, 10
	fmt.Println(e > f) // true

	var g, h int = 20, 20
	fmt.Println(g >= h)

}
