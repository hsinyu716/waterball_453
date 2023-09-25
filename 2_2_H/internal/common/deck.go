package common

import (
	"cosmos.cards.showdown/internal/common/card"
	"math/rand"
)

type Deck struct {
	Cards []card.Card
}

func NewDeck(cards []card.Card) *Deck {
	deck := &Deck{
		cards,
	}
	return deck
}

type IDeck interface {
	Push(card card.Card)
	Shuffle()
	DrawCard() card.Card
	Size() int
	TopCard() card.Card
	IsEmpty() bool
}

func (d *Deck) Push(card0 card.Card) {
	d.Cards = append([]card.Card{card0}, d.Cards...)
}

func (d *Deck) Shuffle() {
	rand.Shuffle(len(d.Cards), func(i, j int) { d.Cards[i], d.Cards[j] = d.Cards[j], d.Cards[i] })
}

func (d *Deck) DrawCard() card.Card {
	card0 := d.Cards[0]
	d.Cards = d.Cards[1:]
	return card0
}

func (d *Deck) Size() int {
	return len(d.Cards)
}

func (d *Deck) TopCard() card.Card {
	card0 := d.Cards[0]
	return card0
}

func (d *Deck) IsEmpty() bool {
	return len(d.Cards) == 0
}
