package main

import "fmt"

// create a new type of 'deck'
// which is a slice of strings
type deck []string //deck extends or borrow behavior of slices of string

// new function that belong to deck type
// function with a receiver, any var of type 'deck' nows get access to the 'print' method
// d is var of cards, in another word it is used for d (conventional way in golang)
// avoid using this or self in golang in referencing a instance of type
// receiver often used abbreviation of type
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
