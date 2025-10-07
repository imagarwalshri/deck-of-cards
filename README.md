# deck-of-cards
Lets build a small game "deck of cards"

## Overview
A simple deck of cards implementation in Go that demonstrates basic card game operations.

## Features
- Create a new deck of 52 cards (4 suits × 13 values)
- Shuffle the deck randomly
- Deal cards from the deck
- Save deck to file
- Load deck from file
- Print deck contents

## Usage

### Build and Run
```bash
go build
./deck-of-cards
```

### Run Tests
```bash
go test
```

## Implementation Details
The deck is implemented as a slice of strings, where each string represents a card (e.g., "Ace of Spades").

The deck includes:
- 4 suits: Spades, Hearts, Diamonds, Clubs
- 13 values: Ace, Two, Three, Four, Five, Six, Seven, Eight, Nine, Ten, Jack, Queen, King
