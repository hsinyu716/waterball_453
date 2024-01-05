package patternHandler

import (
	"cosmos.big2/internal/common/pattern"
	"cosmos.big2/internal/common/poker"
	"fmt"
)

type PairHandler struct {
	IPatternHandler
	nextHandler IPatternHandler
}

func NewPairHandler(pattern IPatternHandler) IPatternHandler {
	return &PairHandler{
		nextHandler: pattern,
	}
}

func (p *PairHandler) Handle(card, topPlay pattern.ICardPattern) int {
	handler := NewPatternHandler(p, p.nextHandler)
	return handler.Handle(card, topPlay)
}

func (p *PairHandler) PrintCard(card pattern.ICardPattern) {
	cardText := "打出了 對子 "
	for _, c := range card.GetCards() {
		cardText = fmt.Sprintf("%s%s[%s] ", cardText, poker.SuitMap[c.GetSuit()], poker.RankMap[c.GetRank()])
	}
	fmt.Println(cardText)
}
