package pattern

import (
	"cosmos.big2/internal/common/poker"
)

type Pair struct {
	CardPattern
}

func NewPatternPair(cards []*poker.Card, next ICardPattern) ICardPattern {
	pair := &Pair{
		CardPattern{
			size:  2,
			cards: cards,
		},
	}
	if pair.Validate() {
		pair.maxCard = *pair.cards[len(pair.cards)-1]
		return pair
	}
	return next
}

func (p *Pair) Validate() bool {
	p.SortRank()
	p.SortSuit()
	if len(p.cards) != p.size {
		return false
	}
	c0 := p.cards[0]
	c1 := p.cards[1]
	return c0.GetRank() == c1.GetRank() && c0.GetSuit() != c1.GetSuit()
}
