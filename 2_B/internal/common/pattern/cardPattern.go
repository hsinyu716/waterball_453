package pattern

import (
	"cosmos.big2/internal/common/poker"
	"fmt"
	"sort"
)

type CardPattern struct {
	size    int
	cards   []*poker.Card
	maxCard poker.Card
}

type ICardPattern interface {
	Validate() bool
	SetCards(cards []*poker.Card)
	GetCards() []*poker.Card
	GetMax() poker.Card
	ShowCard() string
	GetThree() int
}

func (c *CardPattern) Validate() bool {
	return true
}

func (c *CardPattern) SetCards(cards []*poker.Card) {
	c.SortRank()
	//c.SortSuit()
	c.cards = cards
}

func (c *CardPattern) GetCards() []*poker.Card {
	return c.cards
}

func (c *CardPattern) GetMax() poker.Card {
	return c.maxCard
}

func (c *CardPattern) SortRank() {
	sort.Slice(c.cards, func(i, j int) bool {
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
	for _, c := range c.cards {
		cardText = fmt.Sprintf("%s%s[%s] ", cardText, poker.SuitMap[c.GetSuit()], poker.RankMap[c.GetRank()])
	}
	return cardText
}

func (c *CardPattern) GetThree() int {
	return 0
}
