package main

import "fmt"

func modify(a *int) {
	*a += 100
}

// func modify(s *string) {
// 	s = "world"
// }

// func modify(s []int) {
// 	s[0] = 100
//}

// func modify(m map[string]int) {
// 	m["K"] = 75
// }
func main() {
	a := 10
	fmt.Println(a)
	modify(&a)
	fmt.Println(a)

	// a := "hello"
	// fmt.Println(a)
	// modify(a)
	// fmt.Println(a)

	// slice := []int{10, 20, 30}
	// fmt.Println(slice)
	// modify(slice)
	// fmt.Println(slice)

	// ascii_codes := make(map[string]int)
	// ascii_codes["A"] = 65
	// ascii_codes["B"] = 70
	// fmt.Println(ascii_codes)
	// modify(ascii_codes)
	// fmt.Println(ascii_codes)
}
