package main

import "fmt"

func main() {
	// Create a new deck
	cards := newDeck()
	
	fmt.Println("New Deck Created:")
	fmt.Println("Total cards:", len(cards))
	fmt.Println()

	// Shuffle the deck
	cards.shuffle()
	fmt.Println("Deck shuffled!")
	fmt.Println()

	// Deal a hand
	hand, remainingDeck := deal(cards, 5)
	
	fmt.Println("Your hand (5 cards):")
	hand.print()
	fmt.Println()

	fmt.Println("Remaining deck has", len(remainingDeck), "cards")
	fmt.Println()

	// Save deck to file
	fmt.Println("Saving deck to file...")
	err := cards.saveToFile("my_deck")
	if err != nil {
		fmt.Println("Error saving deck:", err)
	} else {
		fmt.Println("Deck saved to 'my_deck' file")
	}
	fmt.Println()

	// Load deck from file
	fmt.Println("Loading deck from file...")
	loadedDeck := newDeckFromFile("my_deck")
	fmt.Println("Loaded deck has", len(loadedDeck), "cards")
}
