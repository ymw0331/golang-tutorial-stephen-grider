package main

import "fmt"

type deck []string

func newDeck() deck { //return a value of type deck
	// Create and return a list of playing cards. Essentially an array of strings
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	// 2 for loop, one nested inside of the other, to iterate all of suits and values
	// replace i and j with _, as never used after declared
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

func print() {
	// Log out the contents of a deck of cards

}

func shuffle() {
	// Shuffles all the cards in a deck
}

func deal() {
	// Create a 'hand' of cards
}

func saveToFile() {
	// Save a list of cards to a file on the local machine

}

func newDeckFromFile() {
	// Load a list of cards from the local machine

}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
