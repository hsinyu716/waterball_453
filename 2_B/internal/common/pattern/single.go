package pattern

import (
	"cosmos.big2/internal/common/poker"
)

type Single struct {
	CardPattern
}

func NewPatternSingle(cards []*poker.Card, next ICardPattern) ICardPattern {
	if len(cards) == 1 {
		return &Single{
			CardPattern{
				size:    1,
				cards:   cards,
				maxCard: *cards[0],
			},
		}
	}
	return next
}
