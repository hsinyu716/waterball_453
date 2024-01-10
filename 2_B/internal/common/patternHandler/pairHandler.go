package patternHandler

import (
	"cosmos.big2/internal/common/pattern"
	"fmt"
	"reflect"
)

type PairHandler struct {
	PatternHandler
	nextHandler IPatternHandler
}

func NewPairHandler(pattern IPatternHandler) IPatternHandler {
	return &PairHandler{
		nextHandler: pattern,
	}
}

func (p *PairHandler) Handle(card, topPlay pattern.ICardPattern) int {
	handler := NewPatternHandler(p, p.nextHandler)
	if reflect.TypeOf(card) != nil && reflect.TypeOf(card).String() == "*pattern.Pair" {
		return handler.Handle(card, topPlay)
	}
	return p.next(p.nextHandler, card, topPlay)
}

func (p *PairHandler) PrintCard(card pattern.ICardPattern) {
	cardText := "打出了 對子 "
	cardString := p.CardString(card)
	fmt.Println(fmt.Sprintf("%s%s", cardText, cardString))
}
