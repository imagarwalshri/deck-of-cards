package main

import (
	"fmt"
	"os"
	"strings"
)

// Creating a new `deck` type as a slice of strings
type deck []string

// newDeck function generates new deck of cards and
// returns a deck instance
func newDeck() deck {
	cards := deck{}

	cardSuites := []string{"Spades", "Clubs", "Diamonds", "Hearts"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuites {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

// print function with deck instance as receiver
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// deal function
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

// toString function: converts deck to string
func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

// save deck to disk
func (d deck) saveToDisk(filename string) error {
	return os.WriteFile(filename, []byte(d.toString()), 0666)
}
