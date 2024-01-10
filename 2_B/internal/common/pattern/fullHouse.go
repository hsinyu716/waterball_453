package pattern

import (
	"cosmos.big2/internal/common/poker"
	"cosmos.big2/internal/utils"
)

type FullHouse struct {
	CardPattern
	nextHandler ICardPattern
	matchThree  int
	size        int
}

func NewPatternFullHouse(next ICardPattern) ICardPattern {
	return &FullHouse{
		nextHandler: next,
		size:        5,
	}
}

func (f *FullHouse) Validate(cards []*poker.Card) ICardPattern {
	if len(cards) == f.size {
		f.SortRank()
		match := map[int]int{}
		var matchNumber []int
		for _, card := range cards {
			if exists, _ := utils.InArray(int(card.GetRank()), matchNumber); !exists {
				matchNumber = append(matchNumber, int(card.GetRank()))
			}
			match[int(card.GetRank())] += 1
			if match[int(card.GetRank())] == 3 {
				f.matchThree = int(card.GetRank())
			}
		}
		isFullHouse := match[matchNumber[0]] == 2 && match[matchNumber[1]] == 3 || match[matchNumber[0]] == 3 && match[matchNumber[1]] == 2
		if isFullHouse {
			f.SetCards(cards)
			f.SortRank()
			f.SetMax(cards[len(cards)-1])
			return f
		}
	}
	if f.nextHandler != nil {
		return f.nextHandler.Validate(cards)
	}
	return nil
}

func (f *FullHouse) GetThree() int {
	return f.matchThree
}
