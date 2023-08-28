package common

import (
	"cosmos.cards.showdown/internal/common/card"
	"fmt"
)

type PlayerAdapter struct {
	name string
	hand *Hand
	game ICardGameStrategy

	PlayerShowdown
}

type IPlayer interface {
	NameHimself(i int) string
	SetHand(hand *Hand)
	AddHandCard(card card.Card)
	GetHand() *Hand
	ShowCard(index int) card.Card
	GetName() string
	SetName(name string)
	GetCardSize() int
	SetGame(game ICardGameStrategy)
	GetGame() ICardGameStrategy

	IPlayerUno
	IPlayerShowdown
}

func (p *PlayerAdapter) NameHimself(i int) string {
	return fmt.Sprintf("AI name %d", i)
}

func (p *PlayerAdapter) SetName(name string) {
	p.name = name
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

func (p *PlayerAdapter) SetGame(game ICardGameStrategy) {
	p.game = game
}

func (p *PlayerAdapter) GetGame() ICardGameStrategy {
	return p.game
}
