package common

import (
	"bufio"
	"cosmos.cards.showdown/internal/common/card"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)

type PlayerAdapter struct {
	name string
	hand *Hand
	game ICardGame

	PlayerShowdown
}

type IPlayer interface {
	NameHimself(i int)
	SetHand(hand *Hand)
	AddHandCard(card card.Card)
	GetHand() *Hand
	ShowCard(index int) card.Card
	GetName() string
	GetCardSize() int
	SetGame(game ICardGame)
	GetGame() ICardGame

	IPlayerUno
	IPlayerShowdown
}

func (p *PlayerAdapter) NameHimself(i int) {
	p.name = fmt.Sprintf("AI random %d", i)
}

func (p *PlayerAdapter) SetHand(hand *Hand) {
	p.hand = hand
}

func (p *PlayerAdapter) AddHandCard(card card.Card) {
	p.hand.AddCard(card)
}

func (p *PlayerAdapter) GetHand() *Hand {
	return p.hand
}

func (p *PlayerAdapter) ShowCard(index int) card.Card {
	return p.GetHand().Show(index)
}

func (p *PlayerAdapter) GetName() string {
	return p.name
}

func (p *PlayerAdapter) GetCardSize() int {
	return len(p.GetHand().Cards)
}

func (p *PlayerAdapter) SetGame(game ICardGame) {
	p.game = game
}

func (p *PlayerAdapter) GetGame() ICardGame {
	return p.game
}
