package pattern

import (
	"cosmos.big2/internal/common/poker"
	"fmt"
	"sort"
)

type ICardPattern interface {
	Validate(cards []*poker.Card) ICardPattern
	SetCards(cards []*poker.Card)
	GetCards() []*poker.Card
	SetMax(card *poker.Card)
	GetMax() poker.Card
	ShowCard() string
	GetThree() int
	SortRank()
	SortSuit()
}

type CardPattern struct {
	size        int
	cards       []*poker.Card
	maxCard     poker.Card
	handler     ICardPattern
	nextHandler ICardPattern
}

func (c *CardPattern) SetCards(cards []*poker.Card) {
	c.SortRank()
	c.cards = cards
}

func (c *CardPattern) GetCards() []*poker.Card {
	return c.cards
}

func (c *CardPattern) SetMax(card *poker.Card) {
	c.maxCard = *card
}

func (c *CardPattern) GetMax() poker.Card {
	return c.maxCard
}

func (c *CardPattern) SortRank() {
	sort.Slice(c.cards, func(i, j int) bool {
		if c.cards[i].GetRank() == c.cards[j].GetRank() {
			return c.cards[i].GetSuit() < c.cards[j].GetSuit()
		}
		return c.cards[i].GetRank() < c.cards[j].GetRank()
	})
}

func (c *CardPattern) SortSuit() {
	sort.Slice(c.cards, func(i, j int) bool {
		return c.cards[i].GetSuit() < c.cards[j].GetSuit()
	})
}

func (c *CardPattern) ShowCard() string {
	cardText := ""
	c.SortRank()
	for _, card := range c.cards {
		cardText = fmt.Sprintf("%s%s[%s] ", cardText, poker.SuitMap[card.GetSuit()], poker.RankMap[card.GetRank()])
	}
	return cardText
}

func (c *CardPattern) GetThree() int {
	return 0
}
