package cards

import (
	"fmt"
	"math/rand"
	"time"
)

type Deck []Card

type Hand Deck

func (d *Deck) Shuffle() {
	// Seed the random number generator to ensure different results each run
	rand.Seed(time.Now().UnixNano())

	// Shuffle the slice
	rand.Shuffle(len(*d), func(i, j int) {
		(*d)[i], (*d)[j] = (*d)[j], (*d)[i]
	})
}

func (d *Deck) Deal() (Hand, Deck) {
	// Deal first 4 cards to each of the players
	hand := (*d)[0:4]
	*d = (*d)[3:]

	return Hand(hand), *d
}

func (h *Hand) Print() {
	// Print the whole deck using fmt
	for _, card := range *h {
		fmt.Printf("%s of %s\n", card.Value, card.Suit)
	}
}

func NewDeck() Deck {
	// Create a new deck of cards
	var deck Deck
	for suit := Hearts; suit <= Spades; suit++ {
		for value := Ace; value <= King; value++ {
			deck = append(deck, Card{suit, value})
		}
	}
	return deck
}
