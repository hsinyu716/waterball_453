package patternHandler

import (
	"cosmos.big2/internal/common/pattern"
	"fmt"
	"reflect"
)

type SingleHandler struct {
	PatternHandler
	handler     IPatternHandler
	nextHandler IPatternHandler
}

func NewSingleHandler(nextHandler IPatternHandler) IPatternHandler {
	return &SingleHandler{
		nextHandler: nextHandler,
	}
}

func (s *SingleHandler) Handle(card, topPlay pattern.ICardPattern) int {
	handler := NewPatternHandler(s, s.nextHandler)
	if reflect.TypeOf(card) != nil && reflect.TypeOf(card).String() == "*pattern.Single" {
		handler.handler = s
		return handler.Handle(card, topPlay)
	}
	return s.next(s.nextHandler, card, topPlay)
}

func (s *SingleHandler) PrintCard(card pattern.ICardPattern) {
	cardText := "打出了 單張 "
	cardString := s.CardString(card)
	fmt.Println(fmt.Sprintf("%s%s", cardText, cardString))
}
