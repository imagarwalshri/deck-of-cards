package main

func main() {
	// Get a new deck of cards
	cards := newDeck()
	// Deal 5 cards
	hand, remainingCards := deal(cards, 5)
	// Print cards in hand
	hand.print()
	// Print remaining cards after deal
	remainingCards.print()
	// Save deck of cards to file
	cards.saveToFile("my_cards")
	// Read deck from file, saved previously
	newCards := newDeckFromFile("my_cards")
	newCards.print()
	// shuffle and deal 5 cards
	shuffledCards := newCards.shuffle()
	shuffledCards.print()
	hand, remainingCards = deal(shuffledCards, 5)
	hand.print()
}
