package common

import (
	"math/rand"
)

type Desk struct {
	cards []*Card
}

func NewDesk() *Desk {
	desk := &Desk{}
	desk.standard52Cards()
	return desk
}

type DeskService interface {
	Push(card *Card)
	Shuffle()
	DrawCard() *Card
	Size() int
}

func (d *Desk) standard52Cards() *Desk {
	suit := []SuitEnumType{Club, Diamond, Heart, Spade}
	rank := []RankEnumType{TWO, THREE, FORE, FIVE, SIX, SEVEN, EIGHT, NINE, TEN, JACK, QUEEN, KING, ACE}
	for _, s := range suit {
		for _, r := range rank {
			d.Push(NewCard(r, s))
		}
	}
	return d
}

func (d *Desk) Push(card *Card) {
	d.cards = append(d.cards, card)
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
