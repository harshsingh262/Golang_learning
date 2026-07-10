package main

import "fmt"

func main() {
	var fruits [5]string = [5]string{"apples", "oranges", "grapes", "mango", "papaya"}
	fmt.Println(fruits[2])
	fmt.Println(fruits)
	fmt.Println(len(fruits))

	marks := [3]int{10, 20, 30}
	fmt.Println(marks)

	names := [...]string{"Rachel", "Phoebe", "Monica"}
	fmt.Println(names)

	grades := [5]int{90, 80, 70, 80, 97}
	fmt.Println(grades)
	grades[1] = 100
	fmt.Println(grades)

	for i := 0; i < len(grades); i++ {
		fmt.Println(grades[i])
	}

	for index, element := range grades {
		fmt.Println(index, "=>", element)
	}

	arr := [3][2]int{{2, 4}, {4, 16}, {8, 64}}
	fmt.Println(arr[2][1])
}
