package main

import "fmt"

func main() {
	for i := 1; i <= 5; i++ {
		if i == 3 {
			continue
		}
		fmt.Println(i * i)
	}

	// x := 1
	// for x <= 5 {
	// 	fmt.Println(x * x)
	// 	x += 1
	// }
}
