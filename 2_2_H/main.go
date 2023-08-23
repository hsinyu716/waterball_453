package main

import (
	"cosmos.cards.showdown/internal/common"
)

func main() {
	p1 := new(common.Human)
	p2 := new(common.AI)
	p3 := new(common.Human)
	p4 := new(common.Human)
	//showdown := common.NewShowdown(&[]common.IPlayer{p1, p2, p3, p4})
	//game0 := common.NewCardGame[*common.GameShowdown](showdown)
	//game0.Start()

	uno := common.NewGameUno(&[]common.IPlayer{p1, p2, p3, p4})
	game2 := common.NewCardGame[*common.CardUno](uno)
	game2.Start()
}
