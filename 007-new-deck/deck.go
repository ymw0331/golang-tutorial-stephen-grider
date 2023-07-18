// Slices are zero-indexed
// fruits = "apple", "banana", "grape", "orange"
// fruits[0] = "apple"
// fruits[3] = "orange"

// Helper to select slices
// fruits[startIndexIncluding : upToNotIncluding]
// fruits[0:2] => "apple", "banana" || or it can be

package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

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

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// Multiple return values
func deal(d deck, handSize int) (deck, deck) {
	//return two values, that return type deck
	//from start of slice to handSize, from handSize to very end
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	//type conversion
	return strings.Join([]string(d), ",")

}

func (d deck) saveToFile(filename string) error {
	// Save a list of cards to a file on the local machine (stdlib writeFile )
	// Converting deck into slice of bytes => deck -> []string -> string -> []byte
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	// Load a list of cards from the local machine
	bs, err := ioutil.ReadFile(filename)
	if err != nil { //error handling
		// Option #1 - log the error and return a call to newDeck()
		// Option #2 - log the error and entirely quit the program
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	s := strings.Split(string(bs), ",")
	// Ace of Spades, Two of Spades, Three of Spades
	return deck(s)

}

func (d deck) shuffle() {

	// Random number generator
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	// Shuffles all the cards in a deck
	// Logic:
	// for each index, card in cards
	// Generate a random number between 0 and len(cards) - 1
	// Swap the current card and the card at cards[randomNumber]

	for i := range d {
		newPosition := r.Intn(len(d) - 1)

		d[i], d[newPosition] = d[newPosition], d[i]
	}

}
