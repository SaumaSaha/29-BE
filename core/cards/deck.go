package cards

import (
	"fmt"
	"math/rand"
	"time"
)

type Deck []Card

func (d *Deck) Shuffle() {
	// Seed the random number generator to ensure different results each run
	rand.Seed(time.Now().UnixNano())

	// Shuffle the slice
	rand.Shuffle(len(*d), func(i, j int) {
		(*d)[i], (*d)[j] = (*d)[j], (*d)[i]
	})
}

func (d *Deck) Print() {
	// Print the whole deck using fmt
	for _, card := range *d {
		fmt.Printf("%s of %s\n", card.Value, card.Suit)
	}
}

func (d *Deck) Deal() (Card, Deck) {
	// Deal first 4-4 cards to each of the players
	card := (*d)[0]
	*d = (*d)[1:]

	return card, *d
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
