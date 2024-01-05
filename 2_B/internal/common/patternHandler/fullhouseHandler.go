package patternHandler

import (
	"cosmos.big2/internal/common/pattern"
	"cosmos.big2/internal/common/poker"
	"fmt"
	"reflect"
)

type FullHouseHandler struct {
	IPatternHandler
	nextHandler IPatternHandler
}

func NewFullHouseHandler(pattern IPatternHandler) IPatternHandler {
	return &FullHouseHandler{
		nextHandler: pattern,
	}
}

func (f *FullHouseHandler) Handle(card, topPlay pattern.ICardPattern) int {
	handler := NewPatternHandler(f, f.nextHandler)
	return handler.Handle(card, topPlay)
}

func (f *FullHouseHandler) Validate(card, topPlay pattern.ICardPattern) bool {
	if card == nil {
		return false
	}
	if card.Validate() && topPlay == nil {
		f.PrintCard(card)
		return true
	}
	if card.Validate() && reflect.TypeOf(card) == reflect.TypeOf(topPlay) {
		topMax := topPlay
		f.PrintCard(card)

		if card.GetThree() > topMax.GetThree() {
			return true
		}
		return false
	}
	return false
}

func (f *FullHouseHandler) PrintCard(card pattern.ICardPattern) {
	cardText := "打出了 葫蘆 "
	for _, c := range card.GetCards() {
		cardText = fmt.Sprintf("%s%s[%s] ", cardText, poker.SuitMap[c.GetSuit()], poker.RankMap[c.GetRank()])
	}
	fmt.Println(cardText)
}
