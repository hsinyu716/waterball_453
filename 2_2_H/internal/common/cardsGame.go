package common

import "fmt"

type CardsGame struct {
	players []IPlayer
	game    interface{}
}

type ICardGame interface {
	nameThemselves()
	setGame()
	Start()
}

func NewCardsGame(players []IPlayer, game interface{}) *CardsGame {
	return &CardsGame{
		players: players,
		game:    game,
	}
}

func (c *CardsGame) nameThemselves() {
	for i, p := range c.players {
		fmt.Println(p)
		p.SetGame(c.game)
		p.NameHimself(i + 1)
		p.SetHand(NewHand(p.GetName()))
	}
}

func (c *CardsGame) Start() {
	//c.game.Start()
}
