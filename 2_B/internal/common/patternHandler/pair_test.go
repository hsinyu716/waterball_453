package patternHandler

import (
	"cosmos.big2/internal/common/pattern"
	"cosmos.big2/internal/common/poker"
	"github.com/stretchr/testify/suite"
	"testing"
)

func TestPairTestSuite(t *testing.T) {
	suite.Run(t, new(PairTestSuite))
}

type PairTestSuite struct {
	suite.Suite
	cards []poker.Card
}

func (s *PairTestSuite) SetupSuite() {
	var cards []poker.Card
	card := poker.NewCard(7, 0)
	cards = append(cards, *card)
	card = poker.NewCard(7, 1)
	cards = append(cards, *card)
	s.cards = cards
}

func (s *PairTestSuite) TestPair() {
	top := pattern.NewPatternPair(s.cards, nil)
	validate := top.Validate()
	s.True(validate)

	var cards []poker.Card
	card := poker.NewCard(9, 1)
	cards = append(cards, *card)
	card = poker.NewCard(9, 2)
	cards = append(cards, *card)
	handler := NewPairHandler(nil)
	handler.Handle(pattern.NewPatternPair(cards, nil), top)
	compare := handler.Validate(pattern.NewPatternPair(cards, nil), top)
	s.True(compare)

	cards = []poker.Card{}
	card = poker.NewCard(7, 2)
	cards = append(cards, *card)
	card = poker.NewCard(7, 3)
	cards = append(cards, *card)
	handler = NewPairHandler(nil)
	compare = handler.Validate(pattern.NewPatternPair(cards, nil), top)
	s.True(compare)
}

func (s *PairTestSuite) TestPairFail() {
	top := pattern.NewPatternPair(s.cards, nil)
	validate := top.Validate()
	s.True(validate)

	var cards []poker.Card
	card := poker.NewCard(3, 1)
	cards = append(cards, *card)
	card = poker.NewCard(3, 2)
	cards = append(cards, *card)
	handler := NewPairHandler(nil)
	compare := handler.Validate(pattern.NewPatternPair(cards, nil), top)
	s.False(compare)

	cards = []poker.Card{}
	card = poker.NewCard(3, 1)
	cards = append(cards, *card)
	handler = NewPairHandler(nil)
	compare = handler.Validate(pattern.NewPatternPair(cards, nil), top)
	s.False(compare)
}
