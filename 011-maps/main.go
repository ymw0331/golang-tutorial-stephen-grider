package main

import "fmt"

// Equivalent of Map in Go
// Map (Go) <-> Hash (Ruby) <-> Object (JS) <-> Dict (Python)

// Both key and value are statically typed of same type

func main() {

	// 1st way of declaring map
	// mapName := make(map[keyType]valueType)
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4b745",
		"white": "#ffffff",
	}

	// 2nd way of declaring map
	// var colors map[string]string

	// 3rd way of declaring map
	// colors := make(map[int]string)

	// In map, we access the value of individual using [] while struct using .
	// colors[10] = "#ffffff"
	// structName.white

	// Built in function
	// delete(colors, 10)

	// fmt.Println(colors)
	printMap(colors)

	// Way to iterate map

}

func printMap(c map[string]string) {
	for color, hex := range c { //declaring and reinitializing all in one step
		fmt.Println("Hex code for", color, "is", hex)

	}
}

// Differences between Struct and Map
// Map
// 1. All key must be the same type
// 2. All values must be same type
// 3. Keys are indexed - we can iterate over them
// 4. Use to represent a collection of related properties
// 5. Don't need to know all the keys at compile time
// 6. Reference Type! 

// Struct
// 1. Value can be of different type
// 2. You nee to know all the different field at compile time
// 3. Keys don't support indexing
// 4. Use to represent a "thing" with a lot of different properties
// 5. Value Type! 
