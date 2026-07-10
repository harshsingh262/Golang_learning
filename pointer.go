package main

import "fmt"

func main() {
	// i := 10
	// fmt.Printf("%T %v \n", &i, &i)
	// fmt.Printf("%T %v \n", *(&i), *(&i))

	// var i *int
	// var s *string
	// fmt.Println(i)
	// fmt.Println(s)

	// s := "hello"
	// var b *string = &s
	// fmt.Println(b)

	// var a = &s
	// fmt.Println(a)

	// c := &s
	// fmt.Println(c)

	s := "Hello"
	fmt.Printf("%T %v \n", s, s)
	ps := &s
	*ps = "world"
	fmt.Printf("%T %v \n", s, s)
	fmt.Printf("%T %v \n", ps, *ps)
}
