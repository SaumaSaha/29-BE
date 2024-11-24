package game

import (
	"29-BE/core/cards"
	"fmt"
)

type Game interface {
	Start()
}

type game struct {
	players []Player
}

func NewGame(players Players) Game {
	return &game{players: players.Players}
}

func (g *game) Start() {
	// Get a new deck
	deck := cards.NewDeck()

	// Shuffle the deck
	deck.Shuffle()

	// Deal the cards
	for i := range g.players {
		hand, remainingDeck := deck.Deal()
		deck = remainingDeck
		g.players[i].Hand = hand
	}

	for _, player := range g.players {
		fmt.Printf("%s's hand:\n", player.Name)
		player.Hand.Print()
	}
}
