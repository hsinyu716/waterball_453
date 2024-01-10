package patternHandler

import (
	"cosmos.big2/internal/common/pattern"
	"fmt"
	"reflect"
)

type FullHouseHandler struct {
	PatternHandler
	nextHandler IPatternHandler
}

func NewFullHouseHandler(pattern IPatternHandler) IPatternHandler {
	return &FullHouseHandler{
		nextHandler: pattern,
	}
}

func (f *FullHouseHandler) Handle(card, topPlay pattern.ICardPattern) int {
	handler := NewPatternHandler(f, f.nextHandler)
	if reflect.TypeOf(card) != nil && reflect.TypeOf(card).String() == "*pattern.FullHouse" {
		return handler.Handle(card, topPlay)
	}
	return f.next(f.nextHandler, card, topPlay)
}

func (f *FullHouseHandler) Validate(card, topPlay pattern.ICardPattern) bool {
	if card == nil {
		return false
	}

	if topPlay == nil {
		f.PrintCard(card)
		return true
	}
	if reflect.TypeOf(card) == reflect.TypeOf(topPlay) {
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
	cardString := f.CardString(card)
	fmt.Println(fmt.Sprintf("%s%s", cardText, cardString))
}
