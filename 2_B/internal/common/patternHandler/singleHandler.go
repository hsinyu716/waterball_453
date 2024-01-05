package patternHandler

import (
	"cosmos.big2/internal/common/pattern"
	"cosmos.big2/internal/common/poker"
	"fmt"
)

type SingleHandler struct {
	IPatternHandler
	nextHandler IPatternHandler
}

func NewSingleHandler(nextHandler IPatternHandler) IPatternHandler {
	return &SingleHandler{
		nextHandler: nextHandler,
	}
}

func (s *SingleHandler) Handle(card, topPlay pattern.ICardPattern) int {
	handler := NewPatternHandler(s, s.nextHandler)
	return handler.Handle(card, topPlay)
}

func (s *SingleHandler) PrintCard(card pattern.ICardPattern) {
	cardText := "打出了 單張 "
	for _, c := range card.GetCards() {
		cardText = fmt.Sprintf("%s%s[%s] ", cardText, poker.SuitMap[c.GetSuit()], poker.RankMap[c.GetRank()])
	}
	fmt.Println(cardText)
}
