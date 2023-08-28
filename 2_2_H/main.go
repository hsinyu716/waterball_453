package main

import (
	"cosmos.cards.showdown/internal/common"
	"cosmos.cards.showdown/internal/common/player"
	"fmt"
)

func main() {
	p1 := new(player.Human)
	p2 := new(player.AI)
	p3 := new(player.Human)
	p4 := new(player.Human)
	//showdown := common.NewShowdown(&[]common.IPlayer{p1, p2, p3, p4})
	//game0 := common.NewCardGame(showdown)
	//game0.Start()

	fmt.Println("===================GAME CHANGE===================")

	uno := common.NewGameUno(&[]common.IPlayer{p1, p2, p3, p4})
	game2 := common.NewCardGame(uno)
	game2.Start()
}
