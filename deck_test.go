package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 52 {
		t.Errorf("Expected deck length of 52, but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card to be Ace of Spades, but got %v", d[0])
	}

	if d[len(d)-1] != "King of Clubs" {
		t.Errorf("Expected last card to be King of Clubs, but got %v", d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	deck := newDeck()
	deck.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 52 {
		t.Errorf("Expected 52 cards in deck, got %v", len(loadedDeck))
	}

	os.Remove("_decktesting")
}

func TestDeal(t *testing.T) {
	d := newDeck()
	hand, remainingDeck := deal(d, 5)

	if len(hand) != 5 {
		t.Errorf("Expected hand size of 5, but got %v", len(hand))
	}

	if len(remainingDeck) != 47 {
		t.Errorf("Expected remaining deck size of 47, but got %v", len(remainingDeck))
	}
}

func TestShuffle(t *testing.T) {
	d := newDeck()
	originalFirst := d[0]
	
	d.shuffle()

	// Check that deck still has 52 cards
	if len(d) != 52 {
		t.Errorf("Expected deck length of 52 after shuffle, but got %v", len(d))
	}

	// It's statistically very unlikely that after shuffling, 
	// the first card remains the same
	// However, this test could occasionally fail due to randomness
	// A better test would be to check that at least some cards changed position
	allSame := true
	originalDeck := newDeck()
	for i := range d {
		if d[i] != originalDeck[i] {
			allSame = false
			break
		}
	}

	if allSame {
		t.Errorf("Shuffle did not change the order of cards")
	}

	// Optional: verify that the first card is different (this could fail occasionally)
	_ = originalFirst // Just to avoid unused variable warning
}
