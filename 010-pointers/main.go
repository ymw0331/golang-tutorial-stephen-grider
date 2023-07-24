package main

import "fmt"

func main() {
	//fmt.Println("Hello, how are you doing")
	mySlice := []string{"Hi", "There", "How", "Are", "You"}

	updateSlice(mySlice)
	fmt.Println(mySlice)
}

// opposite of struct and pointer
func updateSlice(s []string) {
	s[0] = "Bye"
}


// Value Types => int, float, string, bool, structs
// Reference Types => slices, maps, channels, pointers, functions``