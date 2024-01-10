package patternHandler

import (
	"cosmos.big2/internal/common/pattern"
	"cosmos.big2/internal/common/poker"
	"fmt"
	"reflect"
)

type IPatternHandler interface {
	Handle(card, topPlay pattern.ICardPattern) int
	Validate(card, topPlay pattern.ICardPattern) bool
	PrintCard(card pattern.ICardPattern)
	CardString(card pattern.ICardPattern) string
}

type PatternHandler struct {
	handler     IPatternHandler
	nextHandler IPatternHandler
}

func NewPatternHandler(handler, nextHandler IPatternHandler) *PatternHandler {
	return &PatternHandler{
		handler:     handler,
		nextHandler: nextHandler,
	}
}

func (p *PatternHandler) Handle(card, topPlay pattern.ICardPattern) int {
	if p.Validate(card, topPlay) {
		return 1
	} else {
		return p.next(p.nextHandler, card, topPlay)
	}
}

func (p *PatternHandler) Validate(card, topPlay pattern.ICardPattern) bool {
	if card == nil {
		return false
	}
	if topPlay == nil {
		p.handler.PrintCard(card)
		return true
	}
	if reflect.TypeOf(card) == reflect.TypeOf(topPlay) {
		topMax := topPlay.GetMax()
		p.handler.PrintCard(card)
		max := card.GetMax()
		if max.GetRank() > topMax.GetRank() {
			return true
		} else if max.GetRank() == topMax.GetRank() && max.GetSuit() > topMax.GetSuit() {
			return true
		}
		return false
	}
	return false
}

func (p *PatternHandler) next(handler IPatternHandler, card, topPlay pattern.ICardPattern) int {
	if handler != nil {
		return handler.Handle(card, topPlay)
	}
	return 0
}

func (p *PatternHandler) CardString(card pattern.ICardPattern) string {
	cardText := ""
	for _, c := range card.GetCards() {
		cardText = fmt.Sprintf("%s%s[%s] ", cardText, poker.SuitMap[c.GetSuit()], poker.RankMap[c.GetRank()])
	}
	return cardText
}
