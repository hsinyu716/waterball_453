package common

import (
	"math/rand"
)

type Desk struct {
	cards []*Card
}

func NewDesk() *Desk {
	desk := &Desk{}
	return desk
}

type IDesk interface {
	Push(card *Card)
	Shuffle()
	DrawCard() *Card
	Size() int
	TopCard() *Card
}

func (d *Desk) Standard52Cards() *Desk {
	suit := []SuitEnumType{Club, Diamond, Heart, Spade}
	rank := []RankEnumType{TWO, THREE, FORE, FIVE, SIX, SEVEN, EIGHT, NINE, TEN, JACK, QUEEN, KING, ACE}
	for _, s := range suit {
		for _, r := range rank {
			d.Push(NewCard(r, s))
		}
	}
	return d
}

func (d *Desk) Standard5Cards() *Desk {
	color := []ColorEnumType{BLUE, RED, YELLOW, GREEN}
	number := []NumberEnumType{N0, N1, N2, N3, N4, N5, N6, N7, N8, N9}
	for _, s := range color {
		for _, r := range number {
			d.Push(NewCard2(r, s))
		}
	}
	return d
}

func (d *Desk) Push(card *Card) {
	d.cards = append([]*Card{card}, d.cards...)
}

func (d *Desk) Shuffle() {
	rand.Shuffle(len(d.cards), func(i, j int) { d.cards[i], d.cards[j] = d.cards[j], d.cards[i] })
}

func (d *Desk) DrawCard() *Card {
	card := d.cards[0]
	d.cards = d.cards[1:]
	return card
}

func (d *Desk) Size() int {
	return len(d.cards)
}

func (d *Desk) TopCard() *Card {
	card := d.cards[0]
	return card
}
