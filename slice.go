package main

import (
	"fmt"
)

func main() {
	// slice := []int{10, 20, 30}
	// fmt.Println(slice)

	//arr := [10]int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
	// slice_1 := arr[1:8]
	// fmt.Println(slice_1)
	// fmt.Println(len(slice_1))
	// fmt.Println(cap(slice_1))

	// slice := arr[1:8]
	// fmt.Println(slice)

	// sub_slice := slice[0:3]
	// fmt.Println(sub_slice)
	// fmt.Print(cap(sub_slice))

	// slice := make([]int, 5, 8)
	// fmt.Println(slice)
	// fmt.Println(len(slice))
	// fmt.Println(cap(slice))

	// arr := [5]int{10, 20, 30, 40, 50}
	// slice := arr[:3]
	// fmt.Println(arr)
	// fmt.Println(slice)

	// slice[1] = 900
	// fmt.Println("After modification")
	// fmt.Println(arr)
	// fmt.Println(slice)

	// arr := [4]int{10, 20, 30, 40}
	// slice := arr[1:3]
	// fmt.Println(slice)
	// fmt.Println(len(slice))
	// fmt.Println(cap(slice))

	// slice = append(slice, 900, -90, 50)
	// fmt.Println(slice)
	// fmt.Println(len(slice))
	// fmt.Println(cap(slice))

	//arr := [5]int{10, 20, 30, 40, 50}
	// slice := arr[:2]

	// arr2 := [5]int{5, 15, 25, 35, 45}
	// slice2 := arr2[:2]

	// new_slice := append(slice, slice2...)
	// fmt.Println(new_slice)

	// i := 2
	// fmt.Println(arr)
	// slice_1 := arr[:i]
	// slice_2 := arr[i+1:]
	// new_slice := append(slice_1, slice_2...)
	// fmt.Println(new_slice)

	// src_slice := []int{10, 20, 30, 40, 50}
	// dest_slice := make([]int, 3)
	// num := copy(dest_slice, src_slice)
	// fmt.Println(dest_slice)
	// fmt.Println("Number of element copied: ", num)

	arr := []int{10, 20, 30, 40, 50}
	for index, value := range arr {
		fmt.Println(index, "=>", value)
	}

	for _, value := range arr {
		fmt.Println(value)
	}

}
