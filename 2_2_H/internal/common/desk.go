package common

import (
	"math/rand"
)

type Desk struct {
	cards []Card
}

func NewDesk(cards []Card) *Desk {
	desk := &Desk{
		cards,
	}
	return desk
}

type IDesk interface {
	Push(card Card)
	Shuffle()
	DrawCard() Card
	Size() int
	TopCard() Card
}

func (d *Desk) Push(card Card) {
	d.cards = append([]Card{card}, d.cards...)
}

func (d *Desk) Shuffle() {
	rand.Shuffle(len(d.cards), func(i, j int) { d.cards[i], d.cards[j] = d.cards[j], d.cards[i] })
}

func (d *Desk) DrawCard() Card {
	card0 := d.cards[0]
	d.cards = d.cards[1:]
	return card0
}

func (d *Desk) Size() int {
	return len(d.cards)
}

func (d *Desk) TopCard() Card {
	card0 := d.cards[0]
	return card0
}
