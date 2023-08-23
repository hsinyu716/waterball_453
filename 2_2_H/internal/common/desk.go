package common

import (
	"cosmos.cards.showdown/internal/common/card"
	"math/rand"
)

type Desk struct {
	Cards []card.Card
}

func NewDesk(cards []card.Card) *Desk {
	desk := &Desk{
		cards,
	}
	return desk
}

type IDesk interface {
	Push(card card.Card)
	Shuffle()
	DrawCard() card.Card
	Size() int
	TopCard() card.Card
}

func (d *Desk) Push(card0 card.Card) {
	d.Cards = append([]card.Card{card0}, d.Cards...)
}

func (d *Desk) Shuffle() {
	rand.Shuffle(len(d.Cards), func(i, j int) { d.Cards[i], d.Cards[j] = d.Cards[j], d.Cards[i] })
}

func (d *Desk) DrawCard() card.Card {
	card0 := d.Cards[0]
	d.Cards = d.Cards[1:]
	return card0
}

func (d *Desk) Size() int {
	return len(d.Cards)
}

func (d *Desk) TopCard() card.Card {
	card0 := d.Cards[0]
	return card0
}
