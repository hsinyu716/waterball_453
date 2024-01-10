package patternHandler

import (
	"cosmos.big2/internal/common/pattern"
	"fmt"
	"reflect"
)

type StraightHandler struct {
	PatternHandler
	nextHandler IPatternHandler
}

func NewStraightHandler(pattern IPatternHandler) IPatternHandler {
	return &StraightHandler{
		nextHandler: pattern,
	}
}

func (s *StraightHandler) Handle(card, topPlay pattern.ICardPattern) int {
	handler := NewPatternHandler(s, s.nextHandler)
	if reflect.TypeOf(card) != nil && reflect.TypeOf(card).String() == "*pattern.Straight" {
		return handler.Handle(card, topPlay)
	}
	return s.next(s.nextHandler, card, topPlay)
}

func (s *StraightHandler) PrintCard(card pattern.ICardPattern) {
	cardText := "打出了 順子 "
	cardString := s.CardString(card)
	fmt.Println(fmt.Sprintf("%s%s", cardText, cardString))
}
