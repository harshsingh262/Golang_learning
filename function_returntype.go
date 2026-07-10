package main

import "fmt"

// func addNumber(a int, b int) string {
// 	sum := a + b
// 	return sum
// }

// func operation(a int, b int) (int, int) {
// 	sum := a + b
// 	diff := a - b
// 	return sum, diff
// }

func operation(a int, b int) (sum int, diff int) {
	sum = a + b
	diff = a - b
	return
}

func main() {
	sum, difference := operation(20, 10)
	fmt.Println(sum, " ", difference)
}
