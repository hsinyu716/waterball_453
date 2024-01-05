package common

import (
	"cosmos.big2/internal/common/poker"
	"fmt"
	"math/rand"
)

type Deck struct {
	cards []*poker.Card
}

func NewDeck() *Deck {
	deck := &Deck{}
	deck.standard52Cards()
	return deck
}

type IDeck interface {
	Push(card *poker.Card)
	Shuffle()
	DrawCard() *poker.Card
	Size() int
	Reset()
}

func (d *Deck) Reset() {
	d.cards = []*poker.Card{}
}

func (d *Deck) standard52Cards() *Deck {
	suit := []poker.SuitEnumType{poker.Spade, poker.Heart, poker.Diamond, poker.Club}
	rank := []poker.RankEnumType{poker.ACE, poker.TWO, poker.THREE, poker.FORE, poker.FIVE, poker.SIX, poker.SEVEN, poker.EIGHT, poker.NINE, poker.TEN, poker.JACK, poker.QUEEN, poker.KING}
	for _, r := range rank {
		for _, s := range suit {
			d.Push(poker.NewCard(r, s))
		}
	}
	return d
}

func (d *Deck) Push(card *poker.Card) {
	d.cards = append(d.cards, card)
}

func (d *Deck) Shuffle() {
	rand.Shuffle(len(d.cards), func(i, j int) { d.cards[i], d.cards[j] = d.cards[j], d.cards[i] })
}

func (d *Deck) DrawCard() *poker.Card {
	card := d.cards[0]
	d.cards = d.cards[1:]
	return card
}

func (d *Deck) Size() int {
	return len(d.cards)
}

func (d *Deck) ShowCard() {
	cardText := ""
	for _, c := range d.cards {
		cardText = fmt.Sprintf("%s%s[%s] ", cardText, poker.SuitMap[c.GetSuit()], poker.RankMap[c.GetRank()])
	}
	fmt.Println(cardText)
}
