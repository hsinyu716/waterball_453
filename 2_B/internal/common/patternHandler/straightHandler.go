package patternHandler

import (
	"cosmos.big2/internal/common/pattern"
	"cosmos.big2/internal/common/poker"
	"fmt"
)

type StraightHandler struct {
	IPatternHandler
	nextHandler IPatternHandler
}

func NewStraightHandler(pattern IPatternHandler) IPatternHandler {
	return &StraightHandler{
		nextHandler: pattern,
	}
}

func (s *StraightHandler) Handle(card, topPlay pattern.ICardPattern) int {
	handler := NewPatternHandler(s, s.nextHandler)
	return handler.Handle(card, topPlay)
}

func (s *StraightHandler) PrintCard(card pattern.ICardPattern) {
	cardText := "打出了 順子 "
	fmt.Println(card.GetCards())
	for _, c := range card.GetCards() {
		cardText = fmt.Sprintf("%s%s[%s] ", cardText, poker.SuitMap[c.GetSuit()], poker.RankMap[c.GetRank()])
	}
	fmt.Println(cardText)
}
