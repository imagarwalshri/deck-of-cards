package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
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
func (d deck) saveToFile(filename string) error {
	return os.WriteFile(filename, []byte(d.toString()), 0666)
}

// read deck from file
func newDeckFromFile(filename string) deck {
	bs, err := os.ReadFile(filename)
	if err != nil {
		// error handling
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	s := strings.Split(string(bs), ",")

	return deck(s)
}

// shuffle a deck of cards
func (d deck) shuffle() deck {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	l := len(d)

	for i := range d {
		newPosition := r.Intn(l - 1)
		d[i], d[newPosition] = d[newPosition], d[i]
	}
	return d
}
