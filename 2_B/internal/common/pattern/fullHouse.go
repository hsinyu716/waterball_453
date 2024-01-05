package pattern

import (
	"cosmos.big2/internal/common/poker"
	"cosmos.big2/internal/utils"
)

type FullHouse struct {
	CardPattern
	matchThree int
}

func NewPatternFullHouse(cards []*poker.Card, next ICardPattern) ICardPattern {
	fullHouse := &FullHouse{
		CardPattern: CardPattern{
			size:  5,
			cards: cards,
		},
	}
	if fullHouse.Validate() {
		fullHouse.maxCard = *fullHouse.cards[len(fullHouse.cards)-1]
		return fullHouse
	}
	return next
}

func (f *FullHouse) Validate() bool {
	if len(f.cards) != f.size {
		return false
	}
	f.SortRank()
	//f.SortSuit()
	match := map[int]int{}
	var matchNumber []int

	for _, card := range f.cards {
		if exists, _ := utils.InArray(int(card.GetRank()), matchNumber); !exists {
			matchNumber = append(matchNumber, int(card.GetRank()))
		}
		match[int(card.GetRank())] += 1
		if match[int(card.GetRank())] == 3 {
			f.matchThree = int(card.GetRank())
		}
	}
	return match[matchNumber[0]] == 2 && match[matchNumber[1]] == 3 || match[matchNumber[0]] == 3 && match[matchNumber[1]] == 2
}

func (f *FullHouse) GetThree() int {
	return f.matchThree
}
