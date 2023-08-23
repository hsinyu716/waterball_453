package common

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)

type PlayerAdapter[T any] struct {
	name     string
	showdown *Showdown
	hand     *Hand
	game     ICardGame[T]

	PlayerShowdown
}

type IPlayer interface {
	NameHimself(i int)
	SetHand(hand *Hand)
	AddHandCard(card Card)
	GetHand() *Hand
	ShowCard(index int) Card
	GetName() string
	GetCardSize() int
	SetGame(game ICardGame[any])
	GetGame() ICardGame[any]

	IPlayerUno
	IPlayerShowdown
}

func (p *PlayerAdapter[T]) NameHimself(i int) {
	p.name = fmt.Sprintf("AI random %d", i)
}

func (p *PlayerAdapter[T]) SetHand(hand *Hand) {
	p.hand = hand
}

func (p *PlayerAdapter[T]) AddHandCard(card Card) {
	p.hand.AddCard(card)
}

func (p *PlayerAdapter[T]) GetHand() *Hand {
	return p.hand
}

func (p *PlayerAdapter[T]) ShowCard(index int) Card {
	return p.GetHand().Show(index)
}

func (p *PlayerAdapter[T]) GetName() string {
	return p.name
}

func (p *PlayerAdapter[T]) GetCardSize() int {
	return len(p.GetHand().cards)
}

func (p *PlayerAdapter[T]) SetGame(game ICardGame[any]) {
	p.game = game
}

func (p *PlayerAdapter[T]) GetGame() ICardGame[any] {
	return p.game
}
