package pattern

import (
	"cosmos.big2/internal/common/poker"
)

type Single struct {
	CardPattern
	nextHandler ICardPattern
	size        int
}

func NewPatternSingle(next ICardPattern) ICardPattern {
	return &Single{
		nextHandler: next,
		size:        1,
	}
}

func (s *Single) Validate(cards []*poker.Card) ICardPattern {
	if len(cards) == s.size {
		s.SetCards(cards)
		s.SetMax(cards[0])
		return s
	}
	if s.nextHandler != nil {
		return s.nextHandler.Validate(cards)
	}
	return nil
}
