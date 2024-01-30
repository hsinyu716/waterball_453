package pattern

import (
	"cosmos.big2/internal/common/poker"
)

type Straight struct {
	CardPattern
	nextHandler ICardPattern
	size        int
}

func NewPatternStraight(next ICardPattern) ICardPattern {
	return &Straight{
		nextHandler: next,
		size:        5,
	}
}

func (s *Straight) Validate(cards []*poker.Card) ICardPattern {
	if len(cards) == s.size {
		s.SetCards(cards)
		s.SortRank()
		continuous := false
		for i := 1; i < len(cards); i++ {
			continuous = cards[i].DiffRank(cards[i-1]) == 1
			if cards[i].DiffRank(cards[i-1]) == 9 {
				// [{0 1} {1 1} {2 1} {11 1} {12 1}]
				//   3     4     5     A     2
				continuous = true
			} else if i == 3 && cards[i].DiffRank(cards[i-1]) == 9 {
				// [{0 1} {1 1} {2 1} {3 1} {12 1}]
				//   3     4     5     6     2
				continuous = true
			}
			if !continuous {
				break
			}
		}
		if continuous {
			s.SetMax(cards[len(cards)-1])
			return s
		}
	}
	if s.nextHandler != nil {
		return s.nextHandler.Validate(cards)
	}
	return nil
}
