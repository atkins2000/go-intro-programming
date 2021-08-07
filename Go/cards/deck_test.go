package main

import (
	"os"
	"testing"
)

//testing has to be specifically done by _test files not w selenium, mocha or whatever else

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected different deck length of 16 but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card of Ace of Spades, but got %v", d[0])
	}
}

//go has no cleanup in testing section

func TestFileIO(t *testing.T) {
	os.Remove("_decktesting") //removes this file

	d := newDeck()
	d.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 16 {
		t.Errorf("Expected 16 cards in deck but got %v", len(loadedDeck))
	}
	os.Remove("_decktesting")
}
