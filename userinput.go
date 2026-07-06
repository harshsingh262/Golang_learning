package main

import (
	"fmt"
)

func main() {
	/*
		fmt.Println("Enter two numbers:")
		var num1 int
		var num2 int
		fmt.Scanf("%d %d", &num1, &num2)
		sum := num1 + num2
		fmt.Printf("The sum of %d and %d is: %d\n", num1, num2, sum)
		var name string
		fmt.Println("Enter your name: ")
		fmt.Scanf("%s", &name)
		fmt.Printf("Hello, %s! Welcome to Go programming.\n", name)
	*/
	fmt.Println("Enter your name & age: ")
	var name string
	fmt.Scanf("%s", &name)
	var age int
	fmt.Scanf("%d", &age)
	fmt.Printf("Hello, %s! You are %d year old.\n", name, age)

}
