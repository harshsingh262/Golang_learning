package main

import (
	"fmt"
)

func main() {
	city := "London"
	{
		country := "United Kingdom"
		fmt.Printf("City: %s, Country: %s\n", city, country)
	}
	fmt.Printf("City: %s\n", city)
	fmt.Printf("Country: %s\n", country) // This line will cause an error because 'country' is not accessible here
}
