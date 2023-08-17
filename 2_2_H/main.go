package main

import (
	"cosmos.cards.showdown/internal/common"
)

func main() {
	p1 := new(common.Human)
	p2 := new(common.AI)
	p3 := new(common.Human)
	p4 := new(common.Human)

	desk := common.NewDesk()
	desk.Standard52Cards()
	showdown := common.NewShowdown(desk, &[]common.IPlayer{p1, p2, p3, p4})
	game := common.NewCardsGame[*common.Showdown](showdown)
	game.Start()

	//desk = common.NewDesk()
	//desk.Standard5Cards()
	//uno := common.NewUno(desk, &[]common.IPlayer{p1, p2, p3, p4})
	//game2 := common.NewCardsGame[*common.Uno](uno)
	//game2.Start()
}
