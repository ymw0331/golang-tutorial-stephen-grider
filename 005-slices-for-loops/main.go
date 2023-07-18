/*
Go has 2 basic DS for handling list of records (Array and Slices)
1. Array (basic or primitive DS for holding fixed length list of things)
2. Slice (an array that can grow or shrink) ** every element in a slice must be of same type
*/

package main

import "fmt"

func main() {
	// slice of cards declaration
	cards := []string{"Ace of Diamonds", newCard()}
	// add additional card, append return a new slice then assign the var of cards
	cards = append(cards, "Six of Spades")

	// fmt.Println(cards)

	//iterate the slices
	for i, card := range cards {
		fmt.Println(i, card)
	}
	// for forloop, every iteration of cards, throw away the previous index and prev card been declared, so we need to reinitialize

}

func newCard() string {
	return "Five of Diamonds"
}
