package game

type Players struct {
	Players            []Player
	CurrentPlayerIndex int
}

func NewPlayers(players []Player) Players {
	return Players{Players: players, CurrentPlayerIndex: 0}
}

func (p *Players) CurrentPlayer() Player {
	return p.Players[p.CurrentPlayerIndex]
}

func (p *Players) NextPlayer() Player {
	p.CurrentPlayerIndex++
	if p.CurrentPlayerIndex >= len(p.Players) {
		p.CurrentPlayerIndex = 0
	}

	return p.CurrentPlayer()
}
