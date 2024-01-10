package patternHandler

import (
	"cosmos.big2/internal/common/pattern"
	"cosmos.big2/internal/common/poker"
	"github.com/stretchr/testify/suite"
	"testing"
)

func TestPatternTestSuite(t *testing.T) {
	suite.Run(t, new(PatternTestSuite))
}

type PatternTestSuite struct {
	suite.Suite
	cards []*poker.Card
}

func (s *PatternTestSuite) SetupSuite() {
	var cards []*poker.Card
	card := poker.NewCard(7, 0)
	cards = append(cards, card)
	card = poker.NewCard(7, 1)
	cards = append(cards, card)
	s.cards = cards
}

func (s *PatternTestSuite) TestPattern() {
	pair := pattern.NewPatternPair(nil)
	topPlay := pair.Validate(s.cards)
	s.NotNil(topPlay)

	var cards []*poker.Card
	card := poker.NewCard(7, 2)
	cards = append(cards, card)
	card = poker.NewCard(7, 3)
	cards = append(cards, card)
	pair2 := pattern.NewPatternPair(nil)
	cardPattern := pair2.Validate(cards)

	handler := NewSingleHandler(NewPairHandler(NewStraightHandler(nil)))
	compare := handler.Handle(cardPattern, topPlay)
	s.Equal(compare, 1)
}

func (s *PatternTestSuite) TestPatternStraight() {
	var cards []*poker.Card
	card := poker.NewCard(2, 2)
	cards = append(cards, card)
	card = poker.NewCard(3, 3)
	cards = append(cards, card)
	card = poker.NewCard(4, 3)
	cards = append(cards, card)
	card = poker.NewCard(5, 3)
	cards = append(cards, card)
	card = poker.NewCard(6, 3)
	cards = append(cards, card)
	cardPattern := pattern.NewPatternStraight(pattern.NewPatternSingle(nil))
	topPlay := cardPattern.Validate(cards)

	cards = []*poker.Card{}
	card = poker.NewCard(12, 2)
	cards = append(cards, card)
	card = poker.NewCard(0, 3)
	cards = append(cards, card)
	card = poker.NewCard(1, 3)
	cards = append(cards, card)
	card = poker.NewCard(2, 3)
	cards = append(cards, card)
	card = poker.NewCard(3, 3)
	cards = append(cards, card)

	cards = []*poker.Card{}
	card = poker.NewCard(0, 1)
	cards = append(cards, card)
	card = poker.NewCard(1, 1)
	cards = append(cards, card)
	card = poker.NewCard(10, 1)
	cards = append(cards, card)
	card = poker.NewCard(11, 1)
	cards = append(cards, card)
	card = poker.NewCard(12, 1)
	cards = append(cards, card)

	cardPattern2 := pattern.NewPatternStraight(pattern.NewPatternSingle(nil))
	myCard := cardPattern2.Validate(cards)
	handler := NewSingleHandler(NewPairHandler(NewStraightHandler(nil)))
	compare := handler.Handle(myCard, topPlay)
	s.Equal(compare, 1)
}
