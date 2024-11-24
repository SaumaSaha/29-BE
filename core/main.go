package main

import (
	"29-BE/core/game"
	"fmt"
)

func main() {
	fmt.Printf("29 GAME.\n")

	game.NewGame(game.NewPlayers([]game.Player{
		{Name: "Player 1"},
		{Name: "Player 2"},
		{Name: "Player 3"},
		{Name: "Player 4"},
	})).Start()
}
