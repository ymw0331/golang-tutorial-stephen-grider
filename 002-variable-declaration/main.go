package main

import "fmt"

func main() {
	// 1st way of assigning a var
	// var card string = "Ace of Spades"

	// 2nd way of assigning a var
	// alternate way, that tell go that card is a var of type string (:= tell go to figure out the type of data assigned to)
	card := "Ace of Spades"
	card = "Five of Diamonds" //we can just assign a new value after :=, := only for new variable
	fmt.Println(card)

	//new var called card
	//string is telling go compiler to only evaluate type string if only be assigned to the var

	// Go is a static type language
	// Basic Go Types
	// ===============
	// bool => true, false
	// string => "Hi!", "Hows it going?"
	// int => 0, -10000, 99999
	// float64 => 10.00001, 0.00009, -100.003

}
