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
	cards []*poker.Card
}

func (s *PairTestSuite) SetupSuite() {
	var cards []*poker.Card
	card := poker.NewCard(7, 0)
	cards = append(cards, card)
	card = poker.NewCard(7, 1)
	cards = append(cards, card)
	s.cards = cards
}

func (s *PairTestSuite) TestPair() {
	top := pattern.NewPatternPair(nil)
	s.NotNil(top.Validate(s.cards))

	var cards []*poker.Card
	card := poker.NewCard(9, 1)
	cards = append(cards, card)
	card = poker.NewCard(9, 2)
	cards = append(cards, card)
	pair := pattern.NewPatternPair(nil)
	validate := pair.Validate(cards)
	handler := NewPairHandler(nil)
	handle := handler.Handle(validate, top)
	s.T().Log(handle)

	cards = []*poker.Card{}
	card = poker.NewCard(7, 2)
	cards = append(cards, card)
	card = poker.NewCard(7, 3)
	cards = append(cards, card)
	pair.Validate(cards)
	handle = handler.Handle(validate, top)
	s.NotNil(handle)
}

func (s *PairTestSuite) TestPairFail() {
	top := pattern.NewPatternPair(nil)
	topPlay := top.Validate(s.cards)
	s.NotNil(topPlay)

	var cards []*poker.Card
	card := poker.NewCard(3, 1)
	cards = append(cards, card)
	card = poker.NewCard(3, 2)
	cards = append(cards, card)
	handler := NewPairHandler(nil)
	pair := pattern.NewPatternPair(nil)
	cardPattern := pair.Validate(cards)
	handle := handler.Handle(cardPattern, topPlay)
	s.Equal(handle, 0)

	cards = []*poker.Card{}
	card = poker.NewCard(3, 1)
	cards = append(cards, card)
	handler = NewPairHandler(nil)
	pair2 := pattern.NewPatternPair(nil)
	cardPattern = pair2.Validate(cards)
	handle = handler.Handle(cardPattern, topPlay)
	s.Equal(handle, 0)
}
