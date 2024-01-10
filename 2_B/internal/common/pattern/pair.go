package pattern

import (
	"cosmos.big2/internal/common/poker"
	"sort"
)

type Pair struct {
	CardPattern
	nextHandler ICardPattern
	size        int
}

func NewPatternPair(next ICardPattern) ICardPattern {
	return &Pair{
		nextHandler: next,
		size:        2,
	}
}

func (p *Pair) Validate(cards []*poker.Card) ICardPattern {
	sort.Slice(cards, func(i, j int) bool {
		if cards[i].GetRank() == cards[j].GetRank() {
			return cards[i].GetSuit() < cards[j].GetSuit()
		}
		return cards[i].GetRank() < cards[j].GetRank()
	})
	if len(cards) == p.size {
		c0 := cards[0]
		c1 := cards[1]
		isPair := c0.GetRank() == c1.GetRank() && c0.GetSuit() != c1.GetSuit()
		if isPair {
			p.SetCards(cards)
			p.SetMax(cards[len(cards)-1])
			return p
		}
	}
	if p.nextHandler != nil {
		return p.nextHandler.Validate(cards)
	}
	return nil
}
