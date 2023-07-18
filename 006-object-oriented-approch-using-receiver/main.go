package main

func main() {
	// slice of cards declaration
	cards := deck{"Ace of Diamonds", newCard()}
	cards = append(cards, "Six of Spades")

	// //iterate the slices
	// for i, card := range cards {
	// 	fmt.Println(i, card)
	// }

	cards.print()
}

func newCard() string {
	return "Five of Diamonds"
}

// when running the program, make sure both files are called
// command: go run main.go deck.go
