package pattern

import (
	"cosmos.big2/internal/common/poker"
)

type Straight struct {
	CardPattern
}

func NewPatternStraight(cards []*poker.Card, next ICardPattern) ICardPattern {
	straight := &Straight{
		CardPattern{
			size:  5,
			cards: cards,
		},
	}
	if straight.Validate() {
		// todo:: 額外條件 判斷A2345為最小順 同順拿A出來比
		straight.maxCard = *straight.cards[len(straight.cards)-1]
		return straight
	}
	return next
}

func (s *Straight) Validate() bool {
	s.SortRank()
	if len(s.cards) != s.size {
		return false
	}
	continuous := false
	for i := 0; i < len(s.cards)-2; i++ {
		continuous = s.cards[i+1].GetRank()-s.cards[i].GetRank() == 1
		if s.cards[i+1].GetRank()-s.cards[i].GetRank() == 9 {
			// [{0 1} {1 1} {2 1} {11 1} {12 1}]
			//   3     4     5     A     2
			continuous = true
		} else if i == 3 && s.cards[i+1].GetRank()-s.cards[i].GetRank() == 9 {
			// [{0 1} {1 1} {2 1} {3 1} {12 1}]
			//   3     4     5     6     2
			continuous = true
		}
		if !continuous {
			break
		}
	}
	return continuous
}
