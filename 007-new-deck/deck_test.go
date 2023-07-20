package main

import (
	"os"
	"testing"
)

// Test Function Naming Conventions
func TestNewDeck(t *testing.T) { //type of value that is being passed, t is test handler
	d := newDeck()

	if len(d) != 16 { //4 suits x 4 values
		t.Errorf("Expected deck length of 20, but got %v", len(d))
	}

	// if len(d) != 16 {
	// 	t.Error("Expected deck length of 20, but got", len(d))
	// }

	// Test if first card is AOS
	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card of Ace of Spades, but got %v", d[0])
	}

	// Test if last card is FOC
	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Expected last card of Four of Clubs, but got %v", d[len(d)-1])
	}

}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_deckte sting")

	deck := newDeck()
	deck.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 16 {
		t.Errorf("Expected 16 cards in deck, got %v", len(deck))
	}

}
