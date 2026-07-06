package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 42
	var s float64 = float64(x)
	fmt.Printf("The value of s: %v, Type of s: %T\n", s, s)

	var h int = 100
	var p string = strconv.Itoa(h)
	fmt.Printf("The value of p: %v, Type of p: %T\n", p, p)

	var w string = "123"
	i, err := strconv.Atoi(w)
	fmt.Printf("The value of i: %v, Type of i: %T\n", i, i)
	fmt.Printf("The value of err: %v, Type of err: %T\n", err, err)
}
