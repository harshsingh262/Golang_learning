package main

import "fmt"

func main() {
	codes := map[string]string{
		"en": "English",
		"fr": "French",
		"hi": "Hindi"}

	// fmt.Println(codes)
	// delete(codes, "en")
	// fmt.Println(codes)

	// codes["it"] = "Italian"
	// fmt.Println(codes)

	// fmt.Println(codes["en"])
	// fmt.Println(codes["fr"])
	// fmt.Println(codes["hi"])

	// value, found := codes["en"]
	// fmt.Println(found, value)

	// codes := make(map[string]int)
	// fmt.Println(codes)

	// for key, value := range codes {
	// 	fmt.Println(key, "=>", value)
	// }

	for key := range codes {
		delete(codes, key)
	}
	fmt.Println(codes)

}
