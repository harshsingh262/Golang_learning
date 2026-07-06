package main

import (
	"fmt"
	"reflect"
)

func main() {
	var x int = 42
	var s string = "Hello, Go!"
	fmt.Println("Type of x:", reflect.TypeOf(x))
	fmt.Println("Type of s:", reflect.TypeOf(s))

	var f float64 = 3.14
	fmt.Printf("The value of f: %v, Type of f: %T\n", f, f)
}
