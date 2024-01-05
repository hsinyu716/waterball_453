package main

import (
	"cosmos.big2/internal/common"
	"cosmos.big2/internal/common/domain"
)

func main() {
	deck := common.NewDeck()

	players := []domain.IPlayer{
		domain.NewHuman(),
		domain.NewHuman(),
		domain.NewHuman(),
		domain.NewHuman(),
	}

	big2 := domain.NewBig2(players, deck)
	big2.Start()
}
