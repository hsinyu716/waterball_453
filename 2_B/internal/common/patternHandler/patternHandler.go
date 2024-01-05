package patternHandler

import (
	"cosmos.big2/internal/common/pattern"
	"reflect"
)

type IPatternHandler interface {
	Handle(card, topPlay pattern.ICardPattern) int
	Validate(card, topPlay pattern.ICardPattern) bool
	PrintCard(card pattern.ICardPattern)
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
		//p.hand.PlayCard(cards)
		//if h.hand.Size() == 0 {
		//	h.big2.Winner = h
		//	return 1
		//}
		//h.big2.TopPlayer = h
		//h.big2.TopPlay = cardPattern
		//fmt.Println(fmt.Sprintf("目前頂牌玩家為 %s, 頂牌為 %s ", h.big2.TopPlayer.GetName(), h.big2.TopPlay.ShowCard()))
		return 1
	} else {
		return p.next(p.nextHandler, card, topPlay)
	}
}

func (p *PatternHandler) Validate(card, topPlay pattern.ICardPattern) bool {
	if card == nil {
		return false
	}
	if card.Validate() && topPlay == nil {
		p.PrintCard(card)
		return true
	}

	if card.Validate() && reflect.TypeOf(card) == reflect.TypeOf(topPlay) {
		topMax := topPlay.GetMax()
		p.PrintCard(card)
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

func (p *PatternHandler) PrintCard(card pattern.ICardPattern) {
}

func (p *PatternHandler) next(handler IPatternHandler, card, topPlay pattern.ICardPattern) int {
	if handler != nil {
		return handler.Handle(card, topPlay)
	}
	return 0
}
