package main

import "fmt"

// Functions and Returns

func main() {
	card := newCard()
	fmt.Println(card)
}

// every function that returns a value must indicate what type of value it is returning
func newCard() string {
	return "Five of Diamonds"
}


