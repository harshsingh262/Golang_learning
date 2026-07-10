package main

import "fmt"

// type Student struct {
// 	name   string
// 	rollNo int
// marks  []int
// grades map[string]int
//}

type Circle struct {
	x      int
	y      int
	radius int
}

func main() {
	// var s Student
	// fmt.Printf("%+v \n", s)

	// st := new(Student)
	// fmt.Printf("%+v", st)

	// st := Student{
	// 	name:   "Joe",
	// 	rollNo: 12,
	// }

	// fmt.Printf("%+v", st)

	// st := Student{"Joe", 12}
	// fmt.Printf("%+v", st)

	var c Circle
	c.x = 5
	c.y = 5
	c.radius = 5
	fmt.Printf("%+v\n", c)
}
