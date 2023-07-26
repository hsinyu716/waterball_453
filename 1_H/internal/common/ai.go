package common

import "fmt"

type AI struct {
	PlayerAdapter
}

func (h *AI) NameHimself(i int) {
	h.name = fmt.Sprintf("AI random %d", i)
}

func (h *AI) TakeTurn() *TurnMove {
	turnMove := NewTurnMove(&h.PlayerAdapter, nil, nil)
	return turnMove
}
