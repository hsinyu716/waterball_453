package main

import (
	"cosmos.cards.showdown/internal/common"
	"cosmos.cards.showdown/internal/common/player"
	"fmt"
)

func main() {
	sp1 := new(player.Human)
	sp2 := new(player.AI)
	sp3 := new(player.Human)
	sp4 := new(player.Human)
	showdown := common.NewShowdown(&[]common.IPlayer{sp1, sp2, sp3, sp4})
	game0 := common.NewCardGame(showdown)
	game0.Start()

	fmt.Println("===================GAME CHANGE===================")
	p1 := new(player.Human)
	p2 := new(player.AI)
	p3 := new(player.Human)
	p4 := new(player.Human)
	uno := common.NewGameUno(&[]common.IPlayer{p1, p2, p3, p4})
	game2 := common.NewCardGame(uno)
	game2.Start()
}
